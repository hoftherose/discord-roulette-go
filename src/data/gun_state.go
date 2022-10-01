package data

import (
	"math/rand"
	"time"
)

var (
	DefaultChamber        []bool = []bool{}
	DefaultNumChamber     int    = 6
	DefaultNumBullet      int    = 1
	DefaultNumBulletsLeft int    = 1
	DefaultCurrentChamber int    = 0
)

type GunState struct {
	Chambers       []bool `json:"chambers"`
	NumChamber     int    `json:"num_chambers"`
	NumBullets     int    `json:"num_bullets"`
	NumBulletsLeft int    `json:"num_bullets_left"`
	CurrentChamber int    `json:"current_chamber"`
}

var DefaultGunState GunState = GunState{
	DefaultChamber,
	DefaultNumChamber,
	DefaultNumBullet,
	DefaultNumBulletsLeft,
	DefaultCurrentChamber,
}

func (g *GunState) SetNextChamber() {
	g.CurrentChamber = (g.CurrentChamber + 1) % g.NumChamber
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

func (g *GunState) CountBullets(shot bool) {
	if shot {
		g.NumBulletsLeft--
	}
	if g.NumBulletsLeft <= 0 {
		g.SpinChamber()
	}
}
