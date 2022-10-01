package data

import (
	"errors"
	"math/rand"
)

var (
	DefaultLosers      []string = []string{}
	DefaultTurns       []string = []string{}
	DefaultCurrentTurn int      = 0
)

type TableState struct {
	Losers      []string `json:"losers"`
	Turns       []string `json:"turns"`
	CurrentTurn int      `json:"current_turn"`
}

var DefaultTableState TableState = TableState{
	DefaultLosers,
	DefaultTurns,
	DefaultCurrentTurn,
}

func (t *TableState) GetCurrentPlayer() string {
	currPlayer := t.CurrentTurn
	return t.Turns[currPlayer]
}

func (t *TableState) SetNextPlayer() {
	t.CurrentTurn = (t.CurrentTurn + 1) % len(t.Turns)
}

func (t *TableState) SpinTable() {
	rand.Shuffle(len(t.Turns), func(i, j int) { t.Turns[i], t.Turns[j] = t.Turns[j], t.Turns[i] })
}

func (t *TableState) Ongoing() bool {
	return len(t.Turns) > 1
}

func (t *TableState) RemovePlayer(user string) error {
	for i, player := range t.Turns {
		if player != user {
			continue
		}
		t.Turns = append(t.Turns[i+1:], t.Turns[:i]...)
		t.CurrentTurn = 0
		return nil
	}
	return errors.New("Player <@" + user + "> not found")
}
