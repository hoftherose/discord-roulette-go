package data_test

import (
	"fmt"
	"testing"

	"github.com/holy-tech/discord-roulette/src/data"
)

func TestSetNextChamber(t *testing.T) {
	var tests = []struct {
		num_chamber, curr_chamber int
		expected                  int
	}{
		{6, 0, 1},
		{6, 4, 5},
		{6, 5, 0},
		{7, 5, 6},
		{1, 0, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.num_chamber, tt.curr_chamber)
		t.Run(testname, func(t *testing.T) {
			gstate := data.GunState{
				NumChamber:     tt.num_chamber,
				CurrentChamber: tt.curr_chamber,
			}
			gstate.SetNextChamber()
			if gstate.CurrentChamber != tt.expected {
				t.Errorf("got %d, want %d", gstate.CurrentChamber, tt.expected)
			}
		})
	}
}
