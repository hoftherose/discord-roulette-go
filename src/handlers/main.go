package handlers

import (
	"flag"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	i "github.com/holy-tech/discord-roulette/src/interfaces"
)

var (
	readyStatus = "/roulette"
	guildID     = flag.String("guild", "", os.Getenv("GUILD_ID"))
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	session.UpdateGameStatus(0, readyStatus)
}

func AppendHandler(s *discordgo.Session, h *i.Handler) {
	s.AddHandler(func(s *discordgo.Session, ic *discordgo.InteractionCreate) {
		if h.CommandSpecs.Name == ic.ApplicationCommandData().Name {
			h.CommandHandler(s, ic)
		}
	})
	_, err := s.ApplicationCommandCreate(s.State.User.ID, *guildID, h.CommandSpecs)
	if err != nil {
		log.Fatalf("Application command %s failed to load: %v", h.CommandSpecs.Name, err)
	}
}
