package handlers

import (
	"flag"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/holy-tech/discord-roulette/src/data"
)

var (
	readyStatus = "/roulette"
	guildID     = flag.String("guild", "", os.Getenv("GUILD_ID"))
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	session.UpdateGameStatus(0, readyStatus)
}

func AppendHandler(s *discordgo.Session, h *data.Handler) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h.CommandSpecs.Name == i.ApplicationCommandData().Name {
			h.CommandHandler(s, i)
		}
	})
	_, err := s.ApplicationCommandCreate(s.State.User.ID, *guildID, h.CommandSpecs)
	if err != nil {
		log.Fatalf("Application command %s failed to load: %v", h.CommandSpecs.Name, err)
	}
}
