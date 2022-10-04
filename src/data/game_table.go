package data

import (
	"math/rand"
	"time"
)

//go:generate mockgen --destination=./../../mocks/table.go github.com/holy-tech/discord-roulette/src/data Table
type Table interface {
	InitTable(...User)
	SpinTable()
	ShuffleTable()
	AcceptPlayer(string)
	NumPlayers() int
	GetSeating() []User
	SetSeating([]User)
	GetCurrentTurn() int
	SetCurrentTurn(int)
	GetSeed() int64
	SetSeed(int64)
}

var (
	DefaultSeating     []User = []User{}
	DefaultCurrentTurn int    = 0
)

type GameTable struct {
	Seating     []User `json:"turns"`
	CurrentTurn int    `json:"current_turn"`
	Seed        int64
}

var DefaultGameTable *GameTable = &GameTable{
	DefaultSeating,
	DefaultCurrentTurn,
	0,
}

func (t *GameTable) InitTable(players ...User) {
	t.SetSeating(players)
	t.ShuffleTable()
	t.SetCurrentTurn(0)
}

func (t *GameTable) SpinTable() {
	rand.Seed(t.GetSeed())
	newStart := rand.Int() % t.NumPlayers()
	seating := t.GetSeating()
	t.SetSeating(
		append(
			seating[newStart:],
			seating[:newStart]...,
		),
	)
}

func (t *GameTable) ShuffleTable() {
	rand.Seed(t.GetSeed())
	seating := t.GetSeating()
	rand.Shuffle(len(seating), func(i, j int) { seating[i], seating[j] = seating[j], seating[i] })
	t.SetSeating(seating)
}

func (t *GameTable) AcceptPlayer(id string) {
	seating := t.GetSeating()
	for _, user := range seating {
		if user.GetID() == id {
			user.Accept()
		}
	}
}

func (t *GameTable) NumPlayers() int {
	return len(t.GetSeating())
}

func (t *GameTable) GetSeating() []User {
	return t.Seating
}

func (t *GameTable) SetSeating(players []User) {
	t.Seating = players
}

func (t *GameTable) GetCurrentTurn() int {
	return t.CurrentTurn
}

func (t *GameTable) SetCurrentTurn(currentTurn int) {
	t.CurrentTurn = currentTurn
}

func (t *GameTable) GetSeed() int64 {
	if t.Seed != 42 {
		return time.Now().UnixNano()
	}
	return t.Seed
}

func (t *GameTable) SetSeed(seed int64) {
	t.Seed = seed
}
