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
		seed        int64
		expected    []string
	}{
		{[]string{"hello", "world"}, 42, []string{"hello", "world"}},
		{[]string{"1", "2", "3"}, 42, []string{"1", "2", "3"}},
		{[]string{"1", "2", "3", "4", "5"}, 42, []string{"4", "1", "5", "2", "3"}},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("%d", i)
		t.Run(testname, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			table := data.GameTable{}
			players := PopulatePlayersID(mockCtrl, tt.playerNames)
			table.InitTable(players...)
			for i, name := range table.GetSeating() {
				if name.GetID() != tt.expected[i] {
					t.Errorf("got %s, want %s", name.GetID(), tt.expected[i])
				}
			}
		})
	}
}
