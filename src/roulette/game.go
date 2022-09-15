package roulette

import (
	"github.com/bwmarrin/discordgo"
	db "github.com/holy-tech/discord-roulette/src/repo"
)

func ShootTurn(channel string, user *discordgo.User) string {
	accepted, err := db.GameIsAccepted(channel)
	if err != nil {
		return "No shots fired: " + err.Error()
	}
	settings, _ := db.GetGameDocument(channel)
	if !accepted {
		return "Game still is not accepted"
	}
	died, err := settings.Shoot(user)
	if err != nil {
		return "Error: " + err.Error()
	} else if died {
		return "You died <@" + user.ID + ">"
	}
	return "You live <@" + user.ID + ">"
}
