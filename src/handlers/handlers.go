package handlers

import (
	"flag"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	readyStatus = "/roulette"
	guildID     = flag.String("guild", "", os.Getenv("GUILD_ID"))
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	session.UpdateGameStatus(0, readyStatus)
}

func AppendHandler(s *discordgo.Session, h *Handler) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		h.commandHandler(s, i)
	})
	_, err := s.ApplicationCommandCreate(s.State.User.ID, *guildID, h.commandSpecs)
	if err != nil {
		log.Fatalf("Application command %s failed to load: %v", h.commandSpecs.Name, err)
	}
}
