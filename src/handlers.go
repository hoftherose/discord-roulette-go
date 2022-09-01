package handlers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const COMMAND_STRING = "!roulette"

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	session.UpdateGameStatus(0, COMMAND_STRING)
}

func Salute(session *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, COMMAND_STRING) {

		// Find the channel that the message came from.
		c, err := session.State.Channel(m.ChannelID)
		if err != nil {
			return
		}
		if c.Name == "roulette-table" {
			session.ChannelMessageSend(c.ID, fmt.Sprintf("You don't fear death... Type %s while in a voice channel to play a sound.", COMMAND_STRING))
		}

		fmt.Println("You are doing stuff in channel", c)
	}
}
