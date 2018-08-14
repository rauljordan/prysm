package rpc

import (
	"context"
	"fmt"
	"testing"

	"github.com/prysmaticlabs/prysm/beacon-chain/types"
	"github.com/prysmaticlabs/prysm/shared/testutil"
	logTest "github.com/sirupsen/logrus/hooks/test"
)

type mockAnnouncer struct{}

func (m *mockAnnouncer) CanonicalBlockAnnouncement() <-chan *types.Block {
	return make(chan *types.Block)
}

func (m *mockAnnouncer) CanonicalCrystallizedStateAnnouncement() <-chan *types.CrystallizedState {
	return make(chan *types.CrystallizedState)
}

func TestLifecycle(t *testing.T) {
	hook := logTest.NewGlobal()
	rpcService := NewRPCService(context.Background(), &Config{Port: "9999", CertFlag: "alice.crt", KeyFlag: "alice.key"}, &mockAnnouncer{})

	rpcService.Start()

	testutil.AssertLogsContain(t, hook, "Starting service")
	testutil.AssertLogsContain(t, hook, fmt.Sprintf("RPC server listening on port :%s", rpcService.port))

	rpcService.Stop()
	testutil.AssertLogsContain(t, hook, "Stopping service")
}

func TestBadEndpoint(t *testing.T) {
	hook := logTest.NewGlobal()
	rpcService := NewRPCService(context.Background(), &Config{Port: "ralph merkle!!!"}, &mockAnnouncer{})

	rpcService.Start()

	testutil.AssertLogsContain(t, hook, "Starting service")
	testutil.AssertLogsContain(t, hook, fmt.Sprintf("Could not listen to port :%s", rpcService.port))

	rpcService.Stop()
	testutil.AssertLogsContain(t, hook, "Stopping service")
}

func TestInsecureEndpoint(t *testing.T) {
	hook := logTest.NewGlobal()
	rpcService := NewRPCService(context.Background(), &Config{Port: "9999"}, &mockAnnouncer{})

	rpcService.Start()

	testutil.AssertLogsContain(t, hook, "Starting service")
	testutil.AssertLogsContain(t, hook, fmt.Sprintf("RPC server listening on port :%s", rpcService.port))
	testutil.AssertLogsContain(t, hook, "You are using an insecure gRPC connection")

	rpcService.Stop()
	testutil.AssertLogsContain(t, hook, "Stopping service")
}

func TestRPCMethods(t *testing.T) {
	rpcService := NewRPCService(context.Background(), &Config{Port: "9999"}, &mockAnnouncer{})
	if _, err := rpcService.ProposeBlock(context.Background(), nil); err == nil {
		t.Error("Wanted error: unimplemented, received nil")
	}
	if _, err := rpcService.SignBlock(context.Background(), nil); err == nil {
		t.Error("Wanted error: unimplemented, received nil")
	}
}
