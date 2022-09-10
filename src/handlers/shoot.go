package handlers

import (
	"github.com/bwmarrin/discordgo"
	r "github.com/holy-tech/discord-roulette/src/roulette"
)

var ShootHandle = Handler{
	&discordgo.ApplicationCommand{
		Name:                     "roulette-shoot",
		Description:              "Pull the trigger",
		DefaultMemberPermissions: &defaultAdmin,
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: r.GameEnd(i.ChannelID),
			},
		})
	},
}
