package data_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/holy-tech/discord-roulette/mocks"
	"github.com/holy-tech/discord-roulette/src/data"
)

func PopulatePlayersID(ctrl *gomock.Controller, names []string) []data.User {
	list := []data.User{}
	for _, name := range names {
		newUser := mock.NewMockUser(ctrl)
		newUser.EXPECT().GetID().Return(name).AnyTimes()
		list = append(list, newUser)
	}
	return list
}

func TestInitTable(t *testing.T) {
	var tests = []struct {
		playerNames []string
		expected    []string
	}{
		{[]string{"hello", "world"}, []string{"world", "hello"}},
		{[]string{"1", "2", "3"}, []string{"3", "1", "2"}},
		{[]string{"1", "2", "3", "4", "5"}, []string{"3", "4", "5", "1", "2"}},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			table := data.GameTable{}
			players := PopulatePlayersID(mockCtrl, tt.playerNames)
			table.SetSeed(42)
			table.InitTable(players...)
			for i, name := range table.GetSeating() {
				if name.GetID() != tt.expected[i] {
					t.Errorf("got %s, want %s", name.GetID(), tt.expected[i])
				}
			}
		})
	}
}

func TestSpinTable(t *testing.T) {
	var tests = []struct {
		playerNames []string
		expected    []string
	}{
		{[]string{"hello", "world"}, []string{"world", "hello"}},
		{[]string{"1", "2", "3"}, []string{"2", "3", "1"}},
		{[]string{"1", "2", "3", "4", "5"}, []string{"1", "2", "3", "4", "5"}},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			table := data.GameTable{}
			players := PopulatePlayersID(mockCtrl, tt.playerNames)
			table.SetSeed(42)
			table.SetSeating(players)
			table.SpinTable()
			for i, name := range table.GetSeating() {
				if name.GetID() != tt.expected[i] {
					t.Errorf("got %s, want %s", name.GetID(), tt.expected[i])
				}
			}
		})
	}
}

func TestShuffleTable(t *testing.T) {
	var tests = []struct {
		playerNames []string
		expected    []string
	}{
		{[]string{"hello", "world"}, []string{"world", "hello"}},
		{[]string{"1", "2", "3"}, []string{"3", "1", "2"}},
		{[]string{"1", "2", "3", "4", "5"}, []string{"3", "4", "5", "1", "2"}},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			table := data.GameTable{}
			players := PopulatePlayersID(mockCtrl, tt.playerNames)
			table.SetSeed(42)
			table.SetSeating(players)
			table.ShuffleTable()
			for i, name := range table.GetSeating() {
				if name.GetID() != tt.expected[i] {
					t.Errorf("got %s, want %s", name.GetID(), tt.expected[i])
				}
			}
		})
	}
}

func TestNumPlayers(t *testing.T) {
	var tests = []struct {
		playerNames []string
	}{
		{[]string{"hello", "world"}},
		{[]string{"1", "2", "3"}},
		{[]string{"1", "2", "3", "4", "5"}},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			table := data.GameTable{}
			players := PopulatePlayersID(mockCtrl, tt.playerNames)
			table.SetSeating(players)
			if table.NumPlayers() != len(tt.playerNames) {
				t.Errorf("got %d, want %d", table.NumPlayers(), len(tt.playerNames))
			}
		})
	}
}
