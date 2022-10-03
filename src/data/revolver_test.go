package data_test

import (
	"fmt"
	"testing"

	"github.com/holy-tech/discord-roulette/src/data"
)

func TestReloadGun(t *testing.T) {
	var tests = []struct {
		sizeChamber, numBullets int
		expected                []bool
	}{
		{6, 1, []bool{true, false, false, false, false, false}},
		{6, 4, []bool{true, true, true, true, false, false}},
		{7, 2, []bool{true, true, false, false, false, false, false}},
		{3, 3, []bool{true, true, true}},
		{1, 0, []bool{false}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("reload_%d_%d", tt.sizeChamber, tt.numBullets)
		t.Run(testname, func(t *testing.T) {
			r := data.DefaultRevolver
			r.ReloadGun(tt.sizeChamber, tt.numBullets)
			chambers := r.GetChamber()
			if len(chambers) != len(tt.expected) {
				t.Errorf("got chamber of size %d, wanted %d", len(chambers), len(tt.expected))
			}
			for i, chamber := range chambers {
				if chamber != tt.expected[i] {
					t.Errorf("error in chamber %d, got %t wanted %t", i, chamber, tt.expected[i])
				}
			}
		})
	}
}

func TestSpinChamber(t *testing.T) {
	var tests = []struct {
		sizeChamber, numBullets int
		seed                    int64
		expected                []bool
	}{
		{6, 1, 42, []bool{false, false, false, false, false, true}},
		{6, 4, 42, []bool{true, true, true, false, false, true}},
		{7, 2, 42, []bool{false, false, false, false, true, true, false}},
		{3, 3, 42, []bool{true, true, true}},
		{1, 0, 42, []bool{false}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("reload_%d_%d", tt.sizeChamber, tt.numBullets)
		t.Run(testname, func(t *testing.T) {
			r := data.DefaultRevolver
			r.SetSeed(tt.seed)
			r.ReloadGun(tt.sizeChamber, tt.numBullets)
			r.SpinChamber()
			chambers := r.GetChamber()
			if len(chambers) != len(tt.expected) {
				t.Errorf("got chamber of size %d, wanted %d", len(chambers), len(tt.expected))
			}
			for i, chamber := range chambers {
				if chamber != tt.expected[i] {
					t.Errorf("error in chamber %d, got %v wanted %v", i, chambers, tt.expected)
				}
			}
		})
	}
}

func TestShuffleChamber(t *testing.T) {
	var tests = []struct {
		sizeChamber, numBullets int
		seed                    int64
		expected                []bool
	}{
		{6, 1, 42, []bool{false, false, false, false, true, false}},
		{6, 4, 42, []bool{true, true, false, false, true, true}},
		{7, 2, 42, []bool{true, false, false, false, false, true, false}},
		{3, 3, 42, []bool{true, true, true}},
		{1, 0, 42, []bool{false}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("reload_%d_%d", tt.sizeChamber, tt.numBullets)
		t.Run(testname, func(t *testing.T) {
			r := data.DefaultRevolver
			r.SetSeed(tt.seed)
			r.ReloadGun(tt.sizeChamber, tt.numBullets)
			r.ShuffleChamber()
			chambers := r.GetChamber()
			if len(chambers) != len(tt.expected) {
				t.Errorf("got chamber of size %d, wanted %d", len(chambers), len(tt.expected))
			}
			for i, chamber := range chambers {
				if chamber != tt.expected[i] {
					t.Errorf("error in chamber %d, got %v wanted %v", i, chambers, tt.expected)
				}
			}
		})
	}
}
