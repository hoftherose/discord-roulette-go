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
	game, _ := db.GetGameDocument(channel)
	if !accepted {
		return "Game still is not accepted"
	}

	var message string
	shot, err := game.Shoot(user)
	db.UpdateGameDocument(channel, game)
	if err != nil {
		message = "Error: " + err.Error()
	} else if shot {
		message = "You died <@" + user.ID + ">"
	} else {
		message = "You live <@" + user.ID + ">"
	}
	return message + "\nIt is " + db.GetCurrentPlayer(channel) + "turn."
}
