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
	Chamber        []bool `json:"chambers"`
	NumBulletsLeft int    `json:"num_bullets_left"`
	CurrentChamber int    `json:"current_chamber"`
	Seed           int64
}

var DefaultRevolver *Revolver = &Revolver{
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
	r.SetChamber(chamber)
}

func (r *Revolver) SpinChamber() {
	rand.Seed(r.GetSeed())
	chamber := r.GetChamber()
	newStart := rand.Int() % len(chamber)
	r.SetChamber(
		append(
			chamber[newStart:],
			chamber[:newStart]...,
		),
	)
}

func (r *Revolver) ShuffleChamber() {
	rand.Seed(r.GetSeed())
	chamber := r.GetChamber()
	rand.Shuffle(len(chamber), func(i, j int) { chamber[i], chamber[j] = chamber[j], chamber[i] })
	r.SetChamber(chamber)
}

func (r *Revolver) Shoot() bool {
	currChamber := r.GetCurrentChamber()
	chamber := r.GetChamber()
	shot := chamber[currChamber]
	if shot {
		chamber[currChamber] = false
	}
	nextChamber := (r.GetCurrentChamber() + 1) % r.ChamberSize()
	r.SetCurrentChamber(nextChamber)
	return shot
}

func (r *Revolver) GetNumBulletsLeft() int {
	bulletCount := 0
	for _, chamber := range r.GetChamber() {
		if chamber {
			bulletCount += 1
		}
	}
	return bulletCount
}

func (r *Revolver) ChamberSize() int {
	return len(r.GetChamber())
}

func (r *Revolver) GetSeed() int64 {
	if r.Seed != 42 {
		return time.Now().UnixNano()
	}
	return r.Seed
}

func (r *Revolver) SetSeed(seed int64) {
	r.Seed = seed
}

func (r *Revolver) GetCurrentChamber() int {
	return r.CurrentChamber
}

func (r *Revolver) SetCurrentChamber(currentChamber int) {
	r.CurrentChamber = currentChamber
}

func (r *Revolver) GetChamber() []bool {
	return r.Chamber
}

func (r *Revolver) SetChamber(chamber []bool) {
	r.Chamber = chamber
}
