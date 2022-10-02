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
	chamber        []bool `json:"chambers"`
	numBulletsLeft int    `json:"num_bullets_left"`
	currentChamber int    `json:"current_chamber"`
	seed           int64
}

var DefaultRevolver Revolver = Revolver{
	DefaultChamber,
	DefaultNumBulletsLeft,
	DefaultCurrentChamber,
	0,
}

func (r *Revolver) ReloadGun(sizeChamber, numBullets int) {
	chamber := make([]bool, sizeChamber)
	for i := range chamber {
		chamber[i] = i < numBullets
	}
	r.SetSeed(time.Now().UnixNano())
}

func (r *Revolver) SpinChamber() {
	rand.Seed(r.Seed())
	newStart := rand.Int() % len(r.Chamber())
	chamber := r.Chamber()
	r.SetChamber(
		append(
			chamber[newStart:],
			chamber[:newStart]...,
		),
	)
}

func (r *Revolver) ShuffleChamber() {
	chamber := r.Chamber()
	rand.Shuffle(len(chamber), func(i, j int) { chamber[i], chamber[j] = chamber[j], chamber[i] })
	r.SetChamber(chamber)
}

func (r *Revolver) Shoot() bool {
	currChamber := r.CurrentChamber()
	chamber := r.Chamber()
	shot := chamber[currChamber]
	if shot {
		chamber[currChamber] = false
	}
	nextChamber := (r.CurrentChamber() + 1) % r.ChamberSize()
	r.SetCurrentChamber(nextChamber)
	return shot
}

func (r *Revolver) NumBulletsLeft() int {
	bulletCount := 0
	for _, chamber := range r.Chamber() {
		if chamber {
			bulletCount += 1
		}
	}
	return bulletCount
}

func (r *Revolver) ChamberSize() int {
	return len(r.Chamber())
}

func (r *Revolver) Seed() int64 {
	return r.seed
}

func (r *Revolver) SetSeed(seed int64) {
	r.seed = seed
}

func (r *Revolver) CurrentChamber() int {
	return r.currentChamber
}

func (r *Revolver) SetCurrentChamber(currentChamber int) {
	r.currentChamber = currentChamber
}

func (r *Revolver) Chamber() []bool {
	return r.chamber
}

func (r *Revolver) SetChamber(chamber []bool) {
	r.chamber = chamber
}
