package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	session.UpdateGameStatus(0, "!roulette")
}

func Salute(session *discordgo.Session, event *discordgo.MessageCreate) {
	fmt.Println(event.Message)
}
