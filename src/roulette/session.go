package roulette

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	u "github.com/holy-tech/discord-roulette/src"
	d "github.com/holy-tech/discord-roulette/src/data"
	db "github.com/holy-tech/discord-roulette/src/repo"
)

func getOpponentsFromSettings(s *d.GameSettings) []string {
	opponents := make([]string, len(s.Opponents))
	i := 0
	for k := range s.Opponents {
		opponents[i] = "<@" + k + ">"
		i++
	}
	return opponents
}

func GameStart(s *d.GameSettings) string {
	opponents := getOpponentsFromSettings(s)
	resp := fmt.Sprintf(
		`Preparing a %d-shooter with %d bullet(s). Prepare your self: %s`,
		s.GunState.NumChamber, s.GunState.NumBullets, u.JoinStrings(", ", opponents...),
	)
	if err := db.CreateGameDocument(s.Channel, s); err != nil {
		log.Printf("Error creating game document: %v", err)
		resp = fmt.Sprintf("Error: %v", err)
	}
	return resp
}

func ChallengeAccept(channel string, user *discordgo.User) string {
	err := db.AcceptPlayer(channel, user)
	if err != nil {
		return "<@" + user.ID + "> Could not accept: " + err.Error()
	}
	message, ready := db.AwaitingPlayer(channel)
	if ready {
		SetTable(channel)
		message += "\nIt is <@" + db.GetCurrentPlayer(channel) + "> turn."
	}
	return "<@" + user.ID + "> has accepted!!\n" + message
}

func ChallengeDeny(channel string, user *discordgo.User) string {
	resp := "Putting gun away"
	if err := db.DeleteGameDocument(channel); err != nil {
		log.Printf("Error removing game: %v", err)
		resp = fmt.Sprintf("Error: %v", err)
	}
	return "<@" + user.ID + "> has denied!!\n" + resp
}
