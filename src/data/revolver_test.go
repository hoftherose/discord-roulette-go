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
		{6, 1, []bool{false, false, false, false, true, false}},
		{6, 4, []bool{true, true, false, false, true, true}},
		{7, 2, []bool{true, false, false, false, false, true, false}},
		{3, 3, []bool{true, true, true}},
		{1, 0, []bool{false}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("reload_%d_%d", tt.sizeChamber, tt.numBullets)
		t.Run(testname, func(t *testing.T) {
			r := data.DefaultRevolver
			r.SetSeed(42)
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
		expected                []bool
	}{
		{6, 1, []bool{false, false, false, true, false, false}},
		{6, 4, []bool{true, false, false, true, true, true}},
		{7, 2, []bool{false, false, true, false, true, false, false}},
		{3, 3, []bool{true, true, true}},
		{1, 0, []bool{false}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("spin_%d_%d", tt.sizeChamber, tt.numBullets)
		t.Run(testname, func(t *testing.T) {
			r := data.DefaultRevolver
			r.SetSeed(42)
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
		expected                []bool
	}{
		{6, 1, []bool{false, false, true, false, false, false}},
		{6, 4, []bool{true, false, true, true, true, false}},
		{7, 2, []bool{false, false, false, true, false, true, false}},
		{3, 3, []bool{true, true, true}},
		{1, 0, []bool{false}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("shuffle_%d_%d", tt.sizeChamber, tt.numBullets)
		t.Run(testname, func(t *testing.T) {
			r := data.DefaultRevolver
			r.SetSeed(42)
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

func TestShoot(t *testing.T) {
	var tests = []struct {
		sizeChamber, numBullets int
		chambers                []bool
		currChamber             int
		expected                bool
		expectedChamber         int
		expectedNumBullets      int
	}{
		{6, 1, []bool{false, false, false, false, true, false}, 4, true, 5, 0},
		{6, 4, []bool{true, true, false, false, true, true}, 3, false, 4, 4},
		{7, 2, []bool{true, false, false, false, false, true, false}, 6, false, 0, 2},
		{3, 3, []bool{true, true, true}, 1, true, 2, 2},
		{1, 0, []bool{false}, 0, false, 0, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("shoot_%d_%d_%d", tt.sizeChamber, tt.numBullets, tt.currChamber)
		t.Run(testname, func(t *testing.T) {
			r := data.DefaultRevolver
			r.ReloadGun(tt.sizeChamber, tt.numBullets)
			r.SetCurrentChamber(tt.currChamber)
			r.SetChamber(tt.chambers)
			shot := r.Shoot()
			if shot != tt.expected {
				t.Errorf("error shooting, got %t, wanted %t", shot, tt.expected)
			}
			if r.GetCurrentChamber() != tt.expectedChamber {
				t.Errorf("got %d, wanted %d", r.GetCurrentChamber(), tt.expectedChamber)
			}
			if r.GetNumBulletsLeft() != tt.expectedNumBullets {
				t.Errorf("got %d, wanted %d", r.GetCurrentChamber(), tt.expectedChamber)
			}
		})
	}
}

func TestGetNumBulletsLeft(t *testing.T) {
	var tests = []struct {
		chambers []bool
		expected int
	}{
		{[]bool{false, true, false, false, false, false}, 1},
		{[]bool{true, true, false}, 2},
		{[]bool{false, false, false}, 0},
		{[]bool{}, 0},
		{[]bool{true, true, false, true}, 3},
		{[]bool{true, true, true, true, true, true}, 6},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("bullets_left_%d", i)
		t.Run(testname, func(t *testing.T) {
			r := data.DefaultRevolver
			r.SetChamber(tt.chambers)

			if r.GetNumBulletsLeft() != tt.expected {
				t.Errorf("got %d, wanted %d", r.GetCurrentChamber(), tt.expected)
			}
		})
	}
}
