package roulette

import (
	"github.com/bwmarrin/discordgo"
	db "github.com/holy-tech/discord-roulette/src/repo"
)

func Died() bool {
	return true
}

func ShootTurn(channel string, user *discordgo.User) string {
	accepted, err := db.GameIsAccepted(channel)
	if err != nil {
		return "No shots fired: " + err.Error()
	}
	if !accepted {
		return "Game still is not accepted"
	} else if Died() {
		return "You died <@" + user.ID + ">"
	}
	return "You live <@" + user.ID + ">"
}
