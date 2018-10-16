package simulator

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/protobuf/proto"
	"github.com/prysmaticlabs/prysm/beacon-chain/db"
	pb "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	"github.com/prysmaticlabs/prysm/shared/event"
	"github.com/prysmaticlabs/prysm/shared/p2p"
	"github.com/prysmaticlabs/prysm/shared/testutil"
	"github.com/sirupsen/logrus"
	logTest "github.com/sirupsen/logrus/hooks/test"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(ioutil.Discard)
}

type mockP2P struct {
	broadcastHash []byte
}

func (mp *mockP2P) Subscribe(msg proto.Message, channel chan p2p.Message) event.Subscription {
	return new(event.Feed).Subscribe(channel)
}

func (mp *mockP2P) Broadcast(msg proto.Message) {
	mp.broadcastHash = msg.(*pb.BeaconBlockHashAnnounce).GetHash()
}

func (mp *mockP2P) Send(msg proto.Message, peer p2p.Peer) {}

type mockPOWChainService struct{}

func (mpow *mockPOWChainService) LatestBlockHash() common.Hash {
	return common.BytesToHash([]byte{})
}

func setupSimulator(t *testing.T) (*Simulator, *mockP2P) {
	ctx := context.Background()

	config := db.Config{Path: "", Name: "", InMemory: true}
	db, err := db.NewDB(config)
	if err != nil {
		t.Fatalf("could not setup beaconDB: %v", err)
	}

	p2pService := &mockP2P{}

	cfg := &Config{
		BlockRequestBuf: 0,
		P2P:             p2pService,
		Web3Service:     &mockPOWChainService{},
		BeaconDB:        db,
		EnablePOWChain:  true,
	}

	return NewSimulator(ctx, cfg), p2pService
}

func TestLifecycle(t *testing.T) {
	hook := logTest.NewGlobal()
	sim, _ := setupSimulator(t)

	sim.Start()
	testutil.AssertLogsContain(t, hook, "Starting service")
	sim.Stop()
	testutil.AssertLogsContain(t, hook, "Stopping service")

	// The context should have been canceled.
	if sim.ctx.Err() == nil {
		t.Error("context was not canceled")
	}
}

func TestBroadcastBlockHash(t *testing.T) {
	hook := logTest.NewGlobal()
	sim, p2pService := setupSimulator(t)

	slotChan := make(chan uint64)
	requestChan := make(chan p2p.Message)
	exitRoutine := make(chan bool)

	go func() {
		sim.run(slotChan, requestChan)
		<-exitRoutine
	}()

	// trigger a new block
	slotChan <- 1

	// test an invalid block request
	requestChan <- p2p.Message{
		Data: &pb.BeaconBlockRequest{
			Hash: make([]byte, 32),
		},
	}

	// test a valid block request
	blockHash := p2pService.broadcastHash
	requestChan <- p2p.Message{
		Data: &pb.BeaconBlockRequest{
			Hash: blockHash,
		},
	}

	// trigger another block
	slotChan <- 2

	testutil.AssertLogsContain(t, hook, "Broadcast block hash")
	testutil.AssertLogsContain(t, hook, "Requested block not found")
	testutil.AssertLogsContain(t, hook, "Responding to full block request")

	// reset logs
	hook.Reset()

	// ensure that another request for the same block can't be made
	requestChan <- p2p.Message{
		Data: &pb.BeaconBlockRequest{
			Hash: blockHash,
		},
	}

	sim.cancel()
	exitRoutine <- true

	testutil.AssertLogsContain(t, hook, "Requested block not found")
	testutil.AssertLogsDoNotContain(t, hook, "Responding to full block request")

	hook.Reset()
}
