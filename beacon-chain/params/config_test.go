package params

import (
	"testing"
)

func TestValidatorStatusCode(t *testing.T) {
	tests := []struct {
		a ValidatorStatusCode
		b int
	}{
		{a: PendingActivation, b: 0},
		{a: Active, b: 1},
		{a: PendingExit, b: 2},
		{a: PendingWithdraw, b: 3},
		{a: Withdrawn, b: 4},
		{a: Penalized, b: 128},
	}
	for _, tt := range tests {
		if int(tt.a) != tt.b {
			t.Errorf("Incorrect validator status code. Wanted: %d, Got: %d", int(tt.a), tt.b)
		}
	}
}

func TestSpecialRecordTypes(t *testing.T) {
	tests := []struct {
		a SpecialRecordType
		b int
	}{
		{a: Logout, b: 0},
		{a: CasperSlashing, b: 1},
	}
	for _, tt := range tests {
		if int(tt.a) != tt.b {
			t.Errorf("Incorrect special record types. Wanted: %d, Got: %d", int(tt.a), tt.b)
		}
	}
}

func TestValidatorSetDeltaFlags(t *testing.T) {
	tests := []struct {
		a ValidatorSetDeltaFlags
		b int
	}{
		{a: Entry, b: 0},
		{a: Exit, b: 1},
	}
	for _, tt := range tests {
		if int(tt.a) != tt.b {
			t.Errorf("Incorrect validator set delta flags. Wanted: %d, Got: %d", int(tt.a), tt.b)
		}
	}
}
