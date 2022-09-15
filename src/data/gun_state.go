package data

import (
	"math/rand"
	"time"
)

var (
	DefaultLosers      []bool   = []bool{}
	DefaultTurns       []string = []string{}
	DefaultCurrentTurn int      = 0
)

type TableState struct {
	Losers      []bool   `json:"losers"`
	Turns       []string `json:"turns"`
	CurrentTurn int      `json:"current_turn"`
}

var DefaultTableState TableState = TableState{
	DefaultLosers,
	DefaultTurns,
	DefaultCurrentTurn,
}

func (g *GunState) SetNextChamber() {
	curr_chamber := g.CurrentChamber
	g.CurrentChamber = (curr_chamber + 1) % g.NumChamber
}

func (g *GunState) SpinChamber() {
	g.Chambers = make([]bool, g.NumChamber)
	for k := 0; k < g.NumChamber; k++ {
		g.Chambers[k] = k < g.NumBullets
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(g.Chambers), func(i, j int) {
		g.Chambers[i], g.Chambers[j] = g.Chambers[j], g.Chambers[i]
	})
}
