package roulette

import (
	"fmt"
)

const (
	DefaultNumChamber        int64 = 6
	DefaultNumBullet         int64 = 1
	DefaultSpinChamber       bool  = false
	DefaultSpinChamberOnShot bool  = false
	DefaultReplaceBullet     bool  = false
)

type GameSettings struct {
	NumChamber        int64 `default:"6"`
	NumBullet         int64 `default:"1"`
	SpinChamber       bool  `default:"false"`
	SpinChamberOnShot bool  `default:"false"`
	ReplaceBullet     bool  `default:"false"`
}

var DefaultGameSettings GameSettings = GameSettings{
	DefaultNumChamber,
	DefaultNumBullet,
	DefaultSpinChamber,
	DefaultSpinChamberOnShot,
	DefaultReplaceBullet,
}

func GameStart(p *GameSettings) string {
	resp := fmt.Sprintf("Preparing a %d-shooter with %d bullet(s).", p.NumChamber, p.NumBullet)
	return resp
}
