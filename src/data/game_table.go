package data

import (
	"math/rand"
)

var (
	DefaultSeating     []Player = []Player{}
	DefaultCurrentTurn int      = 0
)

type GameTable struct {
	seating     []Player `json:"turns"`
	currentTurn int      `json:"current_turn"`
	seed        int64
}

var DefaultGameTable GameTable = GameTable{
	DefaultSeating,
	DefaultCurrentTurn,
	0,
}

func (t *GameTable) SetTable(players ...Player) {
	t.SetSeating(players)
	t.ShuffleTable()
	t.SetCurrentTurn(0)
}

func (t *GameTable) SpinTable() {
	newStart := rand.Int() % t.NumPlayers()
	seating := t.Seating()
	t.SetSeating(
		append(
			seating[newStart:],
			seating[:newStart]...,
		),
	)
}

func (t *GameTable) ShuffleTable() {
	seating := t.Seating()
	rand.Shuffle(len(seating), func(i, j int) { seating[i], seating[j] = seating[j], seating[i] })
	t.SetSeating(seating)
}

func (t *GameTable) NumPlayers() int {
	return len(t.Seating())
}

func (t *GameTable) Seating() []Player {
	return t.seating
}

func (t *GameTable) SetSeating(players []Player) {
	t.seating = players
}

func (t *GameTable) CurrentTurn() int {
	return t.currentTurn
}

func (t *GameTable) SetCurrentTurn(currentTurn int) {
	t.currentTurn = currentTurn
}

func (t *GameTable) Seed() int64 {
	return t.seed
}

func (t *GameTable) SetSeed(seed int64) {
	t.seed = seed
}
