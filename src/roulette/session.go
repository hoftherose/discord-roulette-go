package roulette

import (
	"fmt"
	"log"

	u "github.com/holy-tech/discord-roulette/src"
	db "github.com/holy-tech/discord-roulette/src/repo"
)

func getOpponentsSettings(s *GameSettings) []string {
	opponents := make([]string, len(s.Opponents))
	for i, o := range s.Opponents {
		opponents[i] = "<@" + o.ID + ">"
	}
	return opponents
}

func GameStart(s *GameSettings) string {
	opponents := getOpponentsSettings(s)
	resp := fmt.Sprintf(
		`Preparing a %d-shooter with %d bullet(s). Prepare your self: %s`,
		s.NumChamber, s.NumBullet, u.JoinStrings(", ", opponents...),
	)
	if err := db.CreateGameDocument(s.Channel, s); err != nil {
		log.Printf("Error creating game document: %v", err)
		resp = fmt.Sprintf("Error: %v", err)
	}
	return resp
}

func ChallengeAccept(s *GameSettings) string {
	return "you accepted!!"
}

func ChallengeDeny(s *GameSettings) string {
	return "you denied!!!"
}

func GameEnd(p *GameSettings) string {
	resp := fmt.Sprintf("Putting gun away\nThe winner is: %s in %s", "Winner", p.Channel)
	if err := db.DeleteGameDocument(p.Channel); err != nil {
		log.Printf("Error removing game: %v", err)
		resp = fmt.Sprintf("Error: %v", err)
	}
	return resp
}
