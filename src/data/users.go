package data

import (
	"github.com/bwmarrin/discordgo"
)

type User struct {
	discordgo.User
	GamesWon    int
	GamesPlayed int
}
