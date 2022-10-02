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
			gstate := data.Revolver{
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

func TestSpinChamber(t *testing.T) {
	var tests = []struct {
		numChamber, numBullets int
		expected               []bool
	}{
		{6, 1, []bool{false, false, false, false, true, false}},
		{6, 3, []bool{true, false, false, false, true, true}},
		{7, 1, []bool{false, false, false, false, false, true, false}},
		{7, 3, []bool{true, false, false, false, false, true, true}},
		{7, 0, []bool{false, false, false, false, false, false, false}},
		{7, 7, []bool{true, true, true, true, true, true, true}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%d", tt.numChamber, tt.numBullets)
		t.Run(testname, func(t *testing.T) {
			gstate := data.Revolver{
				NumBullets: tt.numBullets,
				NumChamber: tt.numChamber,
			}
			gstate.SpinChamber()
			if len(gstate.Chambers) != len(tt.expected) {
				t.Errorf("diff size in got %v, and want %v", gstate.Chambers, tt.expected)
				return
			}
			for i, chamber := range gstate.Chambers {
				if chamber != tt.expected[i] {
					t.Errorf("error in chamber %d got %v, want %v", i, gstate.Chambers, tt.expected)
					return
				}
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
			gstate := data.Revolver{
				NumBulletsLeft: tt.numBulletsLeft,
			}
			gstate.ClearChamber(tt.shot)
			if gstate.NumBulletsLeft != tt.expected {
				t.Errorf("got %d, want %d", gstate.NumBulletsLeft, tt.expected)
			}
		})
	}
}
