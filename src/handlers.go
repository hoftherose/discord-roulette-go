package handlers

import (
	"flag"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	COMMAND_STRING = "!roulette"
	// AppID          = flag.String("app", "", os.Getenv("APP_ID"))
	// GuildID        = flag.String("guild", "", os.Getenv("GUILD_ID"))
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	session.UpdateGameStatus(0, COMMAND_STRING)
}

func AppendHandler(s *discordgo.Session, h *Handler) {
	s.AddHandler(h.commandHandler)
	_, err := s.ApplicationCommandCreate(h.commandSpecs.ApplicationID, h.commandSpecs.GuildID, h.commandSpecs)
	if err != nil {
		log.Fatalf("Application command failed to load: %v", err)
	}
}
