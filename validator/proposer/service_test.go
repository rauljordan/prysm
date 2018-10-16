package proposer

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/golang/mock/gomock"
	pbp2p "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	pb "github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1"
	"github.com/prysmaticlabs/prysm/shared/event"
	"github.com/prysmaticlabs/prysm/shared/testutil"
	"github.com/prysmaticlabs/prysm/validator/internal"
	"github.com/sirupsen/logrus"
	logTest "github.com/sirupsen/logrus/hooks/test"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(ioutil.Discard)
}

type mockClient struct {
	ctrl *gomock.Controller
}

func (mc *mockClient) ProposerServiceClient() pb.ProposerServiceClient {
	return internal.NewMockProposerServiceClient(mc.ctrl)
}

type mockAssigner struct{}

func (m *mockAssigner) ProposerAssignmentFeed() *event.Feed {
	return new(event.Feed)
}

type mockAttesterFeed struct{}

func (m *mockAttesterFeed) ProcessedAttestationFeed() *event.Feed {
	return new(event.Feed)
}

func TestDoesAttestationExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := &Config{
		AssignmentBuf: 0,
		Assigner:      &mockAssigner{},
		Client:        &mockClient{ctrl},
	}
	p := NewProposer(context.Background(), cfg)

	p.pendingAttestation = []*pbp2p.AggregatedAttestation{
		{
			AttesterBitfield: []byte{'a'},
		},
		{
			AttesterBitfield: []byte{'b'},
		},
		{
			AttesterBitfield: []byte{'c'},
		},
		{
			AttesterBitfield: []byte{'d'},
		}}

	fakeAttestation := &pbp2p.AggregatedAttestation{
		AttesterBitfield: []byte{'e'},
	}

	realAttestation := &pbp2p.AggregatedAttestation{
		AttesterBitfield: []byte{'a'},
	}

	if p.DoesAttestationExist(fakeAttestation) {
		t.Fatal("invalid attestation exists")
	}

	if !p.DoesAttestationExist(realAttestation) {
		t.Fatal("valid attestation does not exists")
	}

}
func TestLifecycle(t *testing.T) {
	hook := logTest.NewGlobal()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := &Config{
		AssignmentBuf: 0,
		Assigner:      &mockAssigner{},
		Client:        &mockClient{ctrl},
		AttesterFeed:  &mockAttesterFeed{},
	}
	p := NewProposer(context.Background(), cfg)
	p.Start()
	p.Stop()

	testutil.AssertLogsContain(t, hook, "Starting service")
	testutil.AssertLogsContain(t, hook, "Stopping service")
}

func TestProposerReceiveBeaconBlock(t *testing.T) {
	hook := logTest.NewGlobal()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := &Config{
		AssignmentBuf: 0,
		Assigner:      &mockAssigner{},
		Client:        &mockClient{ctrl},
		AttesterFeed:  &mockAttesterFeed{},
	}
	p := NewProposer(context.Background(), cfg)

	mockServiceClient := internal.NewMockProposerServiceClient(ctrl)
	mockServiceClient.EXPECT().ProposeBlock(
		gomock.Any(),
		gomock.Any(),
	).Return(&pb.ProposeResponse{
		BlockHash: []byte("hi"),
	}, nil)

	doneChan := make(chan struct{})
	exitRoutine := make(chan bool)

	go func() {
		p.run(doneChan, mockServiceClient)
		<-exitRoutine
	}()
	p.assignmentChan <- &pbp2p.BeaconBlock{Slot: 5}
	doneChan <- struct{}{}
	exitRoutine <- true

	testutil.AssertLogsContain(t, hook, "Performing proposer responsibility")
	testutil.AssertLogsContain(t, hook, fmt.Sprintf("Block proposed successfully with hash %#x", []byte("hi")))
	testutil.AssertLogsContain(t, hook, "Proposer context closed")
}

func TestProposerProcessAttestation(t *testing.T) {
	hook := logTest.NewGlobal()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := &Config{
		AssignmentBuf: 0,
		Assigner:      &mockAssigner{},
		Client:        &mockClient{ctrl},
		AttesterFeed:  &mockAttesterFeed{},
	}
	p := NewProposer(context.Background(), cfg)

	doneChan := make(chan struct{})
	exitRoutine := make(chan bool)

	go func() {
		p.processAttestation(doneChan)
		<-exitRoutine
	}()
	p.pendingAttestation = []*pbp2p.AggregatedAttestation{
		{
			AttesterBitfield: []byte{'a'},
		},
		{
			AttesterBitfield: []byte{'b'},
		}}

	attestation := &pbp2p.AggregatedAttestation{AttesterBitfield: []byte{'c'}}
	p.attestationChan <- attestation

	doneChan <- struct{}{}
	exitRoutine <- true

	testutil.AssertLogsContain(t, hook, "Attestation stored in memory")
	testutil.AssertLogsContain(t, hook, "Proposer context closed")

	if !bytes.Equal(p.pendingAttestation[2].GetAttesterBitfield(), []byte{'c'}) {
		t.Errorf("attestation was unable to be saved %v", p.pendingAttestation[2].GetAttesterBitfield())
	}
}

func TestFullProposalOfBlock(t *testing.T) {
	hook := logTest.NewGlobal()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := &Config{
		AssignmentBuf: 0,
		Assigner:      &mockAssigner{},
		Client:        &mockClient{ctrl},
		AttesterFeed:  &mockAttesterFeed{},
	}
	p := NewProposer(context.Background(), cfg)
	mockServiceClient := internal.NewMockProposerServiceClient(ctrl)
	mockServiceClient.EXPECT().ProposeBlock(
		gomock.Any(),
		gomock.Any(),
	).Return(&pb.ProposeResponse{
		BlockHash: []byte("hi"),
	}, nil)

	doneChan := make(chan struct{})
	exitRoutine := make(chan bool)

	go p.run(doneChan, mockServiceClient)

	go func() {
		p.processAttestation(doneChan)
		<-exitRoutine
	}()

	p.pendingAttestation = []*pbp2p.AggregatedAttestation{
		{
			AttesterBitfield: []byte{'a'},
		},
		{
			AttesterBitfield: []byte{'b'},
		}}

	attestation := &pbp2p.AggregatedAttestation{AttesterBitfield: []byte{'c'}}
	p.attestationChan <- attestation

	p.assignmentChan <- &pbp2p.BeaconBlock{Slot: 5}

	doneChan <- struct{}{}
	doneChan <- struct{}{}
	exitRoutine <- true

	testutil.AssertLogsContain(t, hook, "Performing proposer responsibility")
	testutil.AssertLogsContain(t, hook, fmt.Sprintf("Block proposed successfully with hash %#x", []byte("hi")))
	testutil.AssertLogsContain(t, hook, "Proposer context closed")
	testutil.AssertLogsContain(t, hook, "Attestation stored in memory")
	testutil.AssertLogsContain(t, hook, "Proposer context closed")

}

func TestProposerServiceErrors(t *testing.T) {
	hook := logTest.NewGlobal()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := &Config{
		AssignmentBuf: 0,
		Assigner:      &mockAssigner{},
		Client:        &mockClient{ctrl},
		AttesterFeed:  &mockAttesterFeed{},
	}
	p := NewProposer(context.Background(), cfg)

	mockServiceClient := internal.NewMockProposerServiceClient(ctrl)

	// Expect call to throw an error.
	mockServiceClient.EXPECT().ProposeBlock(
		gomock.Any(),
		gomock.Any(),
	).Return(nil, errors.New("bad block proposed"))

	doneChan := make(chan struct{})
	exitRoutine := make(chan bool)

	go p.run(doneChan, mockServiceClient)

	go func() {
		p.processAttestation(doneChan)
		<-exitRoutine
	}()

	p.attestationChan <- &pbp2p.AggregatedAttestation{}
	p.assignmentChan <- nil
	p.assignmentChan <- &pbp2p.BeaconBlock{Slot: 9}

	doneChan <- struct{}{}
	doneChan <- struct{}{}
	exitRoutine <- true

	testutil.AssertLogsContain(t, hook, "Performing proposer responsibility")
	testutil.AssertLogsContain(t, hook, "Could not marshal latest beacon block")
	testutil.AssertLogsContain(t, hook, "Proposer context closed")
	testutil.AssertLogsContain(t, hook, "Could not propose block: bad block proposed")
}
