package data

import (
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
	curr_player := t.CurrentTurn
	return t.Turns[curr_player]
}

func (t *TableState) SetNextPlayer() {
	t.CurrentTurn = (t.CurrentTurn + 1) % len(t.Turns)
}

func (t *TableState) SpinTable() {
	rand.Shuffle(len(t.Turns), func(i, j int) { t.Turns[i], t.Turns[j] = t.Turns[j], t.Turns[i] })
}