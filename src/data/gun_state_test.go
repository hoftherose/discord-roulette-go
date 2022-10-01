package data_test

import (
	"fmt"
	"testing"

	"github.com/holy-tech/discord-roulette/src/data"
)

func TestSetNextChamber(t *testing.T) {
	var tests = []struct {
		numChamber, currChamber int
		expected                int
	}{
		{6, 0, 1},
		{6, 4, 5},
		{6, 5, 0},
		{7, 5, 6},
		{1, 0, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.numChamber, tt.currChamber)
		t.Run(testname, func(t *testing.T) {
			gstate := data.GunState{
				NumChamber:     tt.numChamber,
				CurrentChamber: tt.currChamber,
			}
			gstate.SetNextChamber()
			if gstate.CurrentChamber != tt.expected {
				t.Errorf("got %d, want %d", gstate.CurrentChamber, tt.expected)
			}
		})
	}
}

func TestClearChamber(t *testing.T) {
	var tests = []struct {
		shot           bool
		numBulletsLeft int
		expected       int
	}{
		{true, 1, 0},
		{false, 4, 4},
		{true, 6, 5},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%d", tt.shot, tt.numBulletsLeft)
		t.Run(testname, func(t *testing.T) {
			gstate := data.GunState{
				NumBulletsLeft: tt.numBulletsLeft,
			}
			gstate.ClearChamber(tt.shot)
			if gstate.NumBulletsLeft != tt.expected {
				t.Errorf("got %d, want %d", gstate.NumBulletsLeft, tt.expected)
			}
		})
	}
}
