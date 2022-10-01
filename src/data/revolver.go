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

type Revolver struct {
	Chambers       []bool `json:"chambers"`
	NumChamber     int    `json:"num_chambers"`
	NumBullets     int    `json:"num_bullets"`
	NumBulletsLeft int    `json:"num_bullets_left"`
	CurrentChamber int    `json:"current_chamber"`
}

var DefaultRevolver Revolver = Revolver{
	DefaultChamber,
	DefaultNumChamber,
	DefaultNumBullet,
	DefaultNumBulletsLeft,
	DefaultCurrentChamber,
}

func (r *Revolver) SetNextChamber() {
	r.CurrentChamber = (r.CurrentChamber + 1) % r.NumChamber
}

func (r *Revolver) SpinChamber() {
	rand.Seed(r.getSeed())
	r.Chambers = make([]bool, r.NumChamber)
	for k := 0; k < r.NumChamber; k++ {
		r.Chambers[k] = k < r.NumBullets
	}
	rand.Shuffle(len(r.Chambers), func(i, j int) {
		r.Chambers[i], r.Chambers[j] = r.Chambers[j], r.Chambers[i]
	})
}

func (r *Revolver) ClearChamber(shot bool) {
	if shot {
		r.NumBulletsLeft--
	}
}

func (r *Revolver) getSeed() int64 {
	return time.Now().UnixNano()
}
