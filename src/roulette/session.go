package roulette

import (
	"fmt"
	"log"

	db "github.com/holy-tech/discord-roulette/src/repo"
)

func GameStart(p *GameSettings) string {
	resp := fmt.Sprintf("Preparing a %d-shooter with %d bullet(s).", p.NumChamber, p.NumBullet)
	if err := db.CreateGameDocument(p.Channel, p); err != nil {
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
