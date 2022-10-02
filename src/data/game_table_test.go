package data_test

import (
	"fmt"
	"testing"

	"github.com/holy-tech/discord-roulette/src/data"
)

func TestGetCurrentPlayer(t *testing.T) {
	var tests = []struct {
		currTurn int
		turns    []string
		expected string
	}{
		{0, []string{"hello", "world"}, "hello"},
		{1, []string{"hello", "world"}, "world"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%v", tt.currTurn, tt.turns)
		t.Run(testname, func(t *testing.T) {
			gstate := data.TableState{
				CurrentTurn: tt.currTurn,
				Turns:       tt.turns,
			}
			if gstate.GetCurrentPlayer() != tt.expected {
				t.Errorf("got %s, want %s", gstate.GetCurrentPlayer(), tt.expected)
			}
		})
	}
}
