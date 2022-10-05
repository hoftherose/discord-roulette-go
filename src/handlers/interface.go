package handlers

import (
	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	CommandSpecs   *discordgo.ApplicationCommand
	CommandHandler func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func (h *Handler) GetName() string {
	if h.CommandSpecs == nil {
		return ""
	}
	return h.CommandSpecs.Name
}
