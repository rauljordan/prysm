package rpc

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/prysmaticlabs/prysm/beacon-chain/internal"
	"github.com/prysmaticlabs/prysm/beacon-chain/types"
	pbp2p "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	pb "github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1"
	"github.com/prysmaticlabs/prysm/shared/event"
	"github.com/prysmaticlabs/prysm/shared/testutil"
	"github.com/sirupsen/logrus"
	logTest "github.com/sirupsen/logrus/hooks/test"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(ioutil.Discard)
}

type mockPOWChainService struct{}

func (m *mockPOWChainService) LatestBlockHash() common.Hash {
	return common.BytesToHash([]byte{})
}

type mockAttestationService struct{}

func (m *mockAttestationService) IncomingAttestationFeed() *event.Feed {
	return new(event.Feed)
}

type mockChainService struct {
	blockFeed       *event.Feed
	stateFeed       *event.Feed
	attestationFeed *event.Feed
}

func (m *mockChainService) IncomingBlockFeed() *event.Feed {
	return new(event.Feed)
}

func (m *mockChainService) CanonicalBlockFeed() *event.Feed {
	return m.blockFeed
}

func (m *mockChainService) CanonicalCrystallizedStateFeed() *event.Feed {
	return m.stateFeed
}

func newMockChainService() *mockChainService {
	return &mockChainService{
		blockFeed:       new(event.Feed),
		stateFeed:       new(event.Feed),
		attestationFeed: new(event.Feed),
	}
}

type mockDB struct {
	block   *types.Block
	genesis *types.Block
	cState  *types.CrystallizedState
}

func (m *mockDB) GetCanonicalBlock() (*types.Block, error) {
	return m.block, nil
}

func (m *mockDB) GetCrystallizedState() *types.CrystallizedState {
	return m.cState
}

func (m *mockDB) GetCanonicalBlockForSlot(uint64) (*types.Block, error) {
	return m.genesis, nil
}

func TestLifecycle(t *testing.T) {
	hook := logTest.NewGlobal()
	rpcService := NewRPCService(context.Background(), &Config{
		Port:     "7348",
		CertFlag: "alice.crt",
		KeyFlag:  "alice.key",
	})

	rpcService.Start()

	testutil.AssertLogsContain(t, hook, "Starting service")
	testutil.AssertLogsContain(t, hook, fmt.Sprintf("RPC server listening on port :%s", rpcService.port))

	rpcService.Stop()
	testutil.AssertLogsContain(t, hook, "Stopping service")
}

func TestBadEndpoint(t *testing.T) {
	hook := logTest.NewGlobal()
	rpcService := NewRPCService(context.Background(), &Config{
		Port: "ralph merkle!!!",
	})

	rpcService.Start()

	testutil.AssertLogsContain(t, hook, "Starting service")
	testutil.AssertLogsContain(t, hook, fmt.Sprintf("Could not listen to port :%s", rpcService.port))

	rpcService.Stop()
	testutil.AssertLogsContain(t, hook, "Stopping service")
}

func TestInsecureEndpoint(t *testing.T) {
	hook := logTest.NewGlobal()
	rpcService := NewRPCService(context.Background(), &Config{
		Port: "7777",
	})

	rpcService.Start()

	testutil.AssertLogsContain(t, hook, "Starting service")
	testutil.AssertLogsContain(t, hook, fmt.Sprintf("RPC server listening on port :%s", rpcService.port))
	testutil.AssertLogsContain(t, hook, "You are using an insecure gRPC connection")

	rpcService.Stop()
	testutil.AssertLogsContain(t, hook, "Stopping service")
}

func TestCurrentAssignmentsAndGenesisTime(t *testing.T) {
	mockChain := &mockChainService{}
	mockDB := &mockDB{}
	mockDB.genesis = types.NewGenesisBlock([32]byte{}, [32]byte{})
	var err error
	mockDB.cState, err = types.NewGenesisCrystallizedState(nil)
	if err != nil {
		t.Fatalf("Could not instantiate initial crystallized state: %v", err)
	}
	rpcService := NewRPCService(context.Background(), &Config{
		Port:            "6372",
		BeaconDB:        mockDB,
		ChainService:    mockChain,
		POWChainService: &mockPOWChainService{},
	})

	key := &pb.PublicKey{PublicKey: []byte{}}
	publicKeys := []*pb.PublicKey{key}
	req := &pb.ValidatorAssignmentRequest{
		PublicKeys: publicKeys,
	}

	res, err := rpcService.CurrentAssignmentsAndGenesisTime(context.Background(), req)
	if err != nil {
		t.Errorf("Could not call CurrentAssignments correctly: %v", err)
	}
	genesis := types.NewGenesisBlock([32]byte{}, [32]byte{})
	if res.GenesisTimestamp.String() != genesis.Proto().GetTimestamp().String() {
		t.Errorf("Received different genesis timestamp, wanted: %v, received: %v", genesis.Proto().GetTimestamp(), res.GenesisTimestamp)
	}
}

func TestProposeBlock(t *testing.T) {
	mockChain := &mockChainService{}
	db := &mockDB{}
	db.cState, _ = types.NewGenesisCrystallizedState(nil)
	rpcService := NewRPCService(context.Background(), &Config{
		Port:            "6372",
		ChainService:    mockChain,
		BeaconDB:        db,
		POWChainService: &mockPOWChainService{},
	})
	req := &pb.ProposeRequest{
		SlotNumber: 5,
		ParentHash: []byte("parent-hash"),
		Timestamp:  ptypes.TimestampNow(),
	}
	if _, err := rpcService.ProposeBlock(context.Background(), req); err != nil {
		t.Errorf("Could not propose block correctly: %v", err)
	}
}

func TestAttestHead(t *testing.T) {
	mockChain := &mockChainService{}
	mockAttestationService := &mockAttestationService{}
	rpcService := NewRPCService(context.Background(), &Config{
		Port:               "6372",
		ChainService:       mockChain,
		AttestationService: mockAttestationService,
	})
	req := &pb.AttestRequest{
		Attestation: &pbp2p.AggregatedAttestation{
			Slot:           999,
			Shard:          1,
			ShardBlockHash: []byte{'a'},
		},
	}
	if _, err := rpcService.AttestHead(context.Background(), req); err != nil {
		t.Errorf("Could not attest head correctly: %v", err)
	}
}

func TestLatestAttestationContextClosed(t *testing.T) {
	hook := logTest.NewGlobal()
	mockAttestationService := &mockAttestationService{}
	rpcService := NewRPCService(context.Background(), &Config{
		Port:               "8777",
		SubscriptionBuf:    0,
		AttestationService: mockAttestationService,
	})
	exitRoutine := make(chan bool)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStream := internal.NewMockBeaconService_LatestAttestationServer(ctrl)
	go func(tt *testing.T) {
		if err := rpcService.LatestAttestation(&empty.Empty{}, mockStream); err != nil {
			tt.Errorf("Could not call RPC method: %v", err)
		}
		<-exitRoutine
	}(t)
	rpcService.cancel()
	exitRoutine <- true
	testutil.AssertLogsContain(t, hook, "RPC context closed, exiting goroutine")
}

func TestLatestAttestationFaulty(t *testing.T) {
	attestationService := &mockAttestationService{}
	rpcService := NewRPCService(context.Background(), &Config{
		Port:               "8777",
		SubscriptionBuf:    0,
		AttestationService: attestationService,
	})
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	exitRoutine := make(chan bool)
	attestation := &types.Attestation{}

	mockStream := internal.NewMockBeaconService_LatestAttestationServer(ctrl)
	mockStream.EXPECT().Send(attestation.Proto()).Return(errors.New("something wrong"))
	// Tests a faulty stream.
	go func(tt *testing.T) {
		if err := rpcService.LatestAttestation(&empty.Empty{}, mockStream); err.Error() != "something wrong" {
			tt.Errorf("Faulty stream should throw correct error, wanted 'something wrong', got %v", err)
		}
		<-exitRoutine
	}(t)

	rpcService.incomingAttestation <- attestation
	rpcService.cancel()
	exitRoutine <- true
}

func TestLatestAttestation(t *testing.T) {
	hook := logTest.NewGlobal()
	attestationService := &mockAttestationService{}
	rpcService := NewRPCService(context.Background(), &Config{
		Port:               "8777",
		SubscriptionBuf:    0,
		AttestationService: attestationService,
	})
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	exitRoutine := make(chan bool)
	attestation := &types.Attestation{}
	mockStream := internal.NewMockBeaconService_LatestAttestationServer(ctrl)
	mockStream.EXPECT().Send(attestation.Proto()).Return(nil)
	// Tests a good stream.
	go func(tt *testing.T) {
		if err := rpcService.LatestAttestation(&empty.Empty{}, mockStream); err != nil {
			tt.Errorf("Could not call RPC method: %v", err)
		}
		<-exitRoutine
	}(t)
	rpcService.incomingAttestation <- attestation
	rpcService.cancel()
	exitRoutine <- true

	testutil.AssertLogsContain(t, hook, "Sending attestation to RPC clients")
}

func TestValidatorSlotAndResponsibility(t *testing.T) {
	mockChain := &mockChainService{}
	mockDB := &mockDB{}
	cState, err := types.NewGenesisCrystallizedState(nil)
	if err != nil {
		t.Fatalf("Failed to instantiate genesis state: %v", err)
	}
	mockDB.cState = cState

	rpcService := NewRPCService(context.Background(), &Config{
		Port:         "6372",
		ChainService: mockChain,
		BeaconDB:     mockDB,
	})
	req := &pb.PublicKey{
		PublicKey: []byte{},
	}
	if _, err := rpcService.ValidatorSlotAndResponsibility(context.Background(), req); err != nil {
		t.Errorf("Could not get validator slot: %v", err)
	}
}

func TestValidatorIndex(t *testing.T) {
	mockChain := &mockChainService{}
	mockDB := &mockDB{}
	cState, err := types.NewGenesisCrystallizedState(nil)
	if err != nil {
		t.Fatalf("Failed to instantiate genesis state: %v", err)
	}
	mockDB.cState = cState

	rpcService := NewRPCService(context.Background(), &Config{
		Port:         "6372",
		ChainService: mockChain,
		BeaconDB:     mockDB,
	})
	req := &pb.PublicKey{
		PublicKey: []byte{},
	}
	if _, err := rpcService.ValidatorIndex(context.Background(), req); err != nil {
		t.Errorf("Could not get validator index: %v", err)
	}
}

func TestValidatorShardID(t *testing.T) {
	mockChain := &mockChainService{}
	mockDB := &mockDB{}
	cState, err := types.NewGenesisCrystallizedState(nil)
	if err != nil {
		t.Fatalf("Failed to instantiate genesis state: %v", err)
	}
	mockDB.cState = cState

	rpcService := NewRPCService(context.Background(), &Config{
		Port:         "6372",
		ChainService: mockChain,
		BeaconDB:     mockDB,
	})
	req := &pb.PublicKey{
		PublicKey: []byte{},
	}
	if _, err := rpcService.ValidatorShardID(context.Background(), req); err != nil {
		t.Errorf("Could not get validator shard ID: %v", err)
	}
}

func TestValidatorAssignments(t *testing.T) {
	hook := logTest.NewGlobal()

	mockChain := newMockChainService()
	mockDB := &mockDB{}
	rpcService := NewRPCService(context.Background(), &Config{
		Port:         "6372",
		ChainService: mockChain,
		BeaconDB:     mockDB,
	})

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStream := internal.NewMockBeaconService_ValidatorAssignmentsServer(ctrl)
	mockStream.EXPECT().Send(gomock.Any()).Return(nil)

	key := &pb.PublicKey{PublicKey: []byte{}}
	publicKeys := []*pb.PublicKey{key}
	req := &pb.ValidatorAssignmentRequest{
		PublicKeys: publicKeys,
	}

	exitRoutine := make(chan bool)

	// Tests a validator assignment stream.
	go func(tt *testing.T) {
		if err := rpcService.ValidatorAssignments(req, mockStream); err != nil {
			tt.Errorf("Could not stream validators: %v", err)
		}
		<-exitRoutine
	}(t)

	genesisState, err := types.NewGenesisCrystallizedState(nil)
	if err != nil {
		t.Fatal(err)
	}

	rpcService.canonicalStateChan <- genesisState
	rpcService.cancel()
	exitRoutine <- true
	testutil.AssertLogsContain(t, hook, "Sending new cycle assignments to validator clients")
}
