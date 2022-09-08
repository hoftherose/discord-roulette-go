package roulette

import (
	"fmt"
	"log"

	db "github.com/holy-tech/discord-roulette/src/repo"
)

const (
	DefaultNumChamber        int64  = 6
	DefaultNumBullet         int64  = 1
	DefaultSpinChamber       bool   = false
	DefaultSpinChamberOnShot bool   = false
	DefaultReplaceBullet     bool   = false
	Channel                  string = ""
)

type GameSettings struct {
	NumChamber        int64  `default:"6"`
	NumBullet         int64  `default:"1"`
	SpinChamber       bool   `default:"false"`
	SpinChamberOnShot bool   `default:"false"`
	ReplaceBullet     bool   `default:"false"`
	Channel           string `default:"none"`
}

var DefaultGameSettings GameSettings = GameSettings{
	DefaultNumChamber,
	DefaultNumBullet,
	DefaultSpinChamber,
	DefaultSpinChamberOnShot,
	DefaultReplaceBullet,
	Channel,
}

func GameStart(p *GameSettings) string {
	resp := fmt.Sprintf("Preparing a %d-shooter with %d bullet(s).", p.NumChamber, p.NumBullet)
	if err := db.CreateGameDocument(p.Channel); err != nil {
		log.Printf("Error creating game document: %v", err)
		resp = fmt.Sprintf("Error: %v", err)
	}
	return resp
}

func GameEnd(p *GameSettings) string {
	resp := fmt.Sprintf("Putting gun away\nThe winner is: %s in %s", "Winner", p.Channel)
	if err := db.DeleteGameDocument(p.Channel); err != nil {
		log.Printf("Error removing game: %v", err)
		resp = fmt.Sprintf("Error: %v", err)
	}
	return resp
}
