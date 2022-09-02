package handlers

import (
	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	commandSpecs   *discordgo.ApplicationCommand
	commandHandler func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func (h *Handler) GetName() string {
	if h.commandSpecs == nil {
		return ""
	}
	return h.commandSpecs.Name
}
