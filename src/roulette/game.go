package roulette

import (
	"github.com/bwmarrin/discordgo"
	db "github.com/holy-tech/discord-roulette/src/repo"
)

func GetGameInfo(channel string) string {
	result, _ := db.GetGameDocument(channel)
	return "Info" + result.Channel
}

func ShootTurn(channel string, user *discordgo.User) string {
	var message string
	game, _ := db.GetGameDocument(channel)
	accepted := game.IsAccepted()
	if !accepted {
		return "Game still is not accepted"
	}

	shot, err := game.TakeTurn()
	db.UpdateGameDocument(channel, game)
	if err != nil {
		message = "Error: " + err.Error()
	} else if shot {
		message = "You died " + user.Mention()
	} else {
		message = "You live " + user.Mention()
	}
	if !game.GameFinished() {
		return message + "\nIt is " + game.Table.GetCurrentPlayer().Mention() + " turn."
	}
	db.DeleteGameDocument(channel)
	return message + "\nThe winner is: " + game.Table.GetCurrentPlayer().Mention()
}
