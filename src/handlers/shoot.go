package handlers

import (
	"github.com/bwmarrin/discordgo"
	i "github.com/holy-tech/discord-roulette/src/interfaces"
	r "github.com/holy-tech/discord-roulette/src/roulette"
)

var ShootHandle = i.Handler{
	CommandSpecs: &discordgo.ApplicationCommand{
		Name:                     "roulette-shoot",
		Description:              "Pull the trigger",
		DefaultMemberPermissions: &defaultAdmin,
	},
	CommandHandler: func(s *discordgo.Session, ic *discordgo.InteractionCreate) {
		s.InteractionRespond(ic.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: r.ShootTurn(ic.ChannelID, ic.Member.User),
			},
		})
	},
}
