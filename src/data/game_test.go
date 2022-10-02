package data_test

import (
	"fmt"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/holy-tech/discord-roulette/src/data"
)

func TestShoot(t *testing.T) {
	var tests = []struct {
		currPlayer, currChamber int
		expected                bool
		expectedErr             error
	}{
		{6, 0, true, nil},
		{6, 4, true, nil},
		{6, 5, true, nil},
		{7, 5, true, nil},
		{1, 0, true, nil},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.currPlayer, tt.currChamber)
		t.Run(testname, func(t *testing.T) {
			gsetting := data.GameStatus{}
			shot, err := gsetting.Shoot(&discordgo.User{})
			if shot != tt.expected {
				t.Errorf("got %t, want %t", shot, tt.expected)
			}
			if err.Error() != tt.expectedErr.Error() {
				t.Errorf("got %s, want %s", err.Error(), tt.expectedErr.Error())
			}
		})
	}
}
