package data_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/holy-tech/discord-roulette/mocks"
	"github.com/holy-tech/discord-roulette/src/data"
)

func PopulatePlayersAccept(ctrl *gomock.Controller, accepted []bool) []data.User {
	list := []data.User{}
	for _, accept := range accepted {
		newUser := mock.NewMockUser(ctrl)
		newUser.EXPECT().HasAccepted().Return(accept).AnyTimes()
		list = append(list, newUser)
	}
	return list
}

func TestIsAccepted(t *testing.T) {
	var tests = []struct {
		acceptance []bool
		expected   bool
	}{
		{[]bool{true, false}, false},
		{[]bool{false, false, false}, false},
		{[]bool{true, true}, true},
		{[]bool{true}, true},
		{[]bool{}, true},
	}
	for i, tt := range tests {
		testname := fmt.Sprintf("game_finished_%d", i)
		t.Run(testname, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			table := mock.NewMockTable(mockCtrl)
			players := PopulatePlayersAccept(
				mockCtrl,
				tt.acceptance,
			)
			table.EXPECT().
				GetSeating().
				Return(
					players,
				)
			gun := mock.NewMockGun(mockCtrl)
			game := data.GameStatus{table, gun, false, ""}
			finished := game.IsAccepted()
			if finished != tt.expected {
				t.Errorf("got %t, want %t", finished, tt.expected)
			}
		})
	}
}

func TestGameFinished(t *testing.T) {
	var tests = []struct {
		numPlayers int
		expected   bool
	}{
		{6, false},
		{3, false},
		{1, true},
		{0, true},
		{-1, true},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("game_finished_%d", tt.numPlayers)
		t.Run(testname, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			table := mock.NewMockTable(mockCtrl)
			table.EXPECT().NumPlayers().Return(tt.numPlayers)
			gun := mock.NewMockGun(mockCtrl)
			game := data.GameStatus{table, gun, false, ""}
			finished := game.GameFinished()
			if finished != tt.expected {
				t.Errorf("got %t, want %t", finished, tt.expected)
			}
		})
	}
}
