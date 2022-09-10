package roulette

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
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

func ChallengeAccept(user *discordgo.User, channel string) string {
	return "<@" + user.ID + "> has accepted!!"
}

func ChallengeDeny(user *discordgo.User, channel string) string {
	resp := GameEnd(channel)
	return "<@" + user.ID + "> has denied!!\n" + resp
}

func GameEnd(channel string) string {
	resp := fmt.Sprintf("Putting gun away\nThe winner is: %s in %s", "Winner", channel)
	if err := db.DeleteGameDocument(channel); err != nil {
		log.Printf("Error removing game: %v", err)
		resp = fmt.Sprintf("Error: %v", err)
	}
	return resp
}
