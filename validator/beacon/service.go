package beacon

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum/event"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	pbp2p "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	pb "github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/blake2b"
)

var log = logrus.WithField("prefix", "beacon")

type rpcClientService interface {
	BeaconServiceClient() pb.BeaconServiceClient
}

// Service that interacts with a beacon node via RPC.
type Service struct {
	ctx                    context.Context
	cancel                 context.CancelFunc
	rpcClient              rpcClientService
	validatorIndex         int
	assignedSlot           uint64
	responsibility         string
	attesterAssignmentFeed *event.Feed
	proposerAssignmentFeed *event.Feed
}

// NewBeaconValidator instantiates a service that interacts with a beacon node.
func NewBeaconValidator(ctx context.Context, rpcClient rpcClientService) *Service {
	ctx, cancel := context.WithCancel(ctx)
	return &Service{
		ctx:                    ctx,
		cancel:                 cancel,
		rpcClient:              rpcClient,
		attesterAssignmentFeed: new(event.Feed),
		proposerAssignmentFeed: new(event.Feed),
	}
}

// Start the main routine for a beacon service.
func (s *Service) Start() {
	log.Info("Starting service")
	client := s.rpcClient.BeaconServiceClient()

	// First thing the validator does is request the current, canonical
	// beacon chain head and crystallized state in case the beacon node is
	// currently in a cycle we are supposed to participate in.
	res, err := client.CanonicalHeadAndState(s.ctx, &empty.Empty{})
	if err != nil {
		// If this RPC request fails, the entire system should fatal as it is critical for
		// the validator to begin this way.
		log.Fatalf("Could not fetch current canonical head and state from beacon node: %v", err)
	}
	if err := s.processCrystallizedState(res.GetCrystallizedState(), client); err != nil {
		log.Fatalf("Unable to process received crystallized state: %v", err)
	}
	// We process the received canonical block to determine if we have to perform
	// validator responsibilities right away.
	s.processBeaconBlock(res.GetCanonicalBlock())

	// Then, we start up some routines in the background that will constantly listen
	// for the canonical head/state in order to keep performing validator responsibilities
	// forever.
	go s.listenForBeaconBlocks(client)
	go s.listenForCrystallizedStates(client)
}

// Stop the main loop..
func (s *Service) Stop() error {
	defer s.cancel()
	log.Info("Stopping service")
	return nil
}

// AttesterAssignmentFeed returns a feed that is written to whenever it is the validator's
// slot to perform attestations.
func (s *Service) AttesterAssignmentFeed() *event.Feed {
	return s.attesterAssignmentFeed
}

// ProposerAssignmentFeed returns a feed that is written to whenever it is the validator's
// slot to proposer blocks.
func (s *Service) ProposerAssignmentFeed() *event.Feed {
	return s.proposerAssignmentFeed
}

// listenForBeaconBlocks receives latest canonical blocks from the beacon node's RPC
// server via gRPC streams.
func (s *Service) listenForBeaconBlocks(client pb.BeaconServiceClient) {
	stream, err := client.LatestBeaconBlock(s.ctx, &empty.Empty{})
	if err != nil {
		log.Errorf("Could not setup beacon chain block streaming client: %v", err)
		return
	}
	for {
		block, err := stream.Recv()
		// If the stream is closed, we stop the loop.
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Errorf("Could not receive latest beacon block from stream: %v", err)
			continue
		}
		s.processBeaconBlock(block)
	}
}

// processBeaconBlock verifies if the slot number matches the validator's assigned
// responsibility. If it does, it notifies the attester/proposer running service
// that it is time to perform.
func (s *Service) processBeaconBlock(block *pbp2p.BeaconBlock) {
	log.WithField("slotNumber", block.GetSlotNumber()).Info("Latest beacon block slot number")

	// Based on the slot determined from the latest crystallized state, check if
	// it matches the latest received beacon slot.
	if s.responsibility == "proposer" {
		log.WithField("slotNumber", block.GetSlotNumber()).Info("Assigned proposal slot number reached")
		s.responsibility = ""
		s.proposerAssignmentFeed.Send(block)
	} else if s.responsibility == "attester" && block.GetSlotNumber() == s.assignedSlot {
		// TODO: Let the validator know a few slots in advance if its attestation slot is coming up
		log.Info("Assigned attestation slot number reached")
		s.responsibility = ""
		s.attesterAssignmentFeed.Send(block)
	}
	return
}

// listenForCrystallizedStates receives the latest canonical crystallized state
// from the beacon node's RPC server via gRPC streams.
func (s *Service) listenForCrystallizedStates(client pb.BeaconServiceClient) {
	stream, err := client.LatestCrystallizedState(s.ctx, &empty.Empty{})
	if err != nil {
		log.Errorf("Could not setup crystallized beacon state streaming client: %v", err)
		return
	}
	for {
		crystallizedState, err := stream.Recv()

		// If the stream is closed, we stop the loop.
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Errorf("Could not receive latest crystallized beacon state from stream: %v", err)
			continue
		}
		if err := s.processCrystallizedState(crystallizedState, client); err != nil {
			log.Error(err)
		}
	}
}

// processCrystallizedState uses a received crystallized state to determine
// whether a validator is a proposer/attester and the validator's assigned slot.
func (s *Service) processCrystallizedState(crystallizedState *pbp2p.CrystallizedState, client pb.BeaconServiceClient) error {
	var activeValidatorIndices []int
	// After receiving the crystallized state, get its hash, and
	// this attester's index in the list.
	stateData, err := proto.Marshal(crystallizedState)
	if err != nil {
		return fmt.Errorf("Could not marshal crystallized state proto: %v", err)
	}

	var crystallizedStateHash [32]byte
	h := blake2b.Sum512(stateData)
	copy(crystallizedStateHash[:], h[:32])

	dynasty := crystallizedState.GetCurrentDynasty()

	for i, validator := range crystallizedState.GetValidators() {
		if validator.StartDynasty <= dynasty && dynasty < validator.EndDynasty {
			activeValidatorIndices = append(activeValidatorIndices, i)
		}
	}
	isValidatorIndexSet := false
	for _, val := range activeValidatorIndices {
		// TODO: Check the public key instead of withdrawal address. This will use BLS.
		if isZeroAddress(crystallizedState.Validators[val].WithdrawalAddress) {
			s.validatorIndex = val
			isValidatorIndexSet = true
			break
		}
	}

	// If validator was not found in the validator set was not set, keep listening for
	// crystallized states.
	if !isValidatorIndexSet {
		log.Debug("Validator index not found in latest crystallized state's active validator list")
		return nil
	}

	req := &pb.ShuffleRequest{
		CrystallizedStateHash: crystallizedStateHash[:],
	}

	res, err := client.FetchShuffledValidatorIndices(s.ctx, req)
	if err != nil {
		return fmt.Errorf("Could not fetch shuffled validator indices: %v", err)
	}

	shuffledIndices := res.GetShuffledValidatorIndices()
	if uint64(s.validatorIndex) == shuffledIndices[len(shuffledIndices)-1] {
		// The validator needs to propose the next block.
		s.responsibility = "proposer"
		log.Debug("Validator selected as proposer of the next slot")
		return nil
	}

	// If the condition above did not pass, the validator is an attester.
	s.responsibility = "attester"

	// Based on the cutoff and assigned slots, determine the beacon block
	// slot at which attester has to perform its responsibility.
	currentAssignedSlots := res.GetAssignedAttestationSlots()
	currentCutoffs := res.GetCutoffIndices()

	// The algorithm functions as follows:
	// Given a list of slots: [0 19 38 57 12 31 50] and
	// A list of cutoff indices: [0 142 285 428 571 714 857 1000]
	// if the validator index is between 0-142, it can attest at slot 0, if it is
	// between 142-285, that validator can attest at slot 19, etc.
	slotIndex := 0
	for i := 0; i < len(currentCutoffs)-1; i++ {
		lowCutoff := currentCutoffs[i]
		highCutoff := currentCutoffs[i+1]
		if (uint64(s.validatorIndex) >= lowCutoff) && (uint64(s.validatorIndex) <= highCutoff) {
			break
		}
		slotIndex++
	}
	s.assignedSlot = currentAssignedSlots[slotIndex]
	log.Debug("Validator selected as attester at slot number: %d", s.assignedSlot)
	return nil
}

// isZeroAddress compares a withdrawal address to an empty byte array.
func isZeroAddress(withdrawalAddress []byte) bool {
	return bytes.Equal(withdrawalAddress, []byte{})
}
