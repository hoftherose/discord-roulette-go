package data

import (
	"math/rand"
)

var (
	DefaultChamber        []bool = []bool{}
	DefaultNumChamber     int    = 6
	DefaultNumBullet      int    = 1
	DefaultCurrentChamber int    = 0
)

type GunState struct {
	Chambers       []bool `json:"chambers"`
	NumChamber     int    `json:"num_chambers"`
	NumBullets     int    `json:"num_bullets"`
	CurrentChamber int    `json:"current_chamber"`
}

var DefaultGunState GunState = GunState{
	DefaultChamber,
	DefaultNumChamber,
	DefaultNumBullet,
	DefaultCurrentChamber,
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
