package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var defaultAdmin int64 = discordgo.PermissionAdministrator

var RouletteHandle = Handler{
	&discordgo.ApplicationCommand{
		Name:                     "roulette-command",
		Description:              "Roulette command",
		DefaultMemberPermissions: &defaultAdmin,
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		resp := fmt.Sprintf("Hey there! Congratulations, you just executed your first slash command in")
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: resp,
			},
		})
	},
}
