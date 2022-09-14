package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/holy-tech/discord-roulette/src/data"
	r "github.com/holy-tech/discord-roulette/src/roulette"
)

var ShootHandle = data.Handler{
	CommandSpecs: &discordgo.ApplicationCommand{
		Name:                     "roulette-shoot",
		Description:              "Pull the trigger",
		DefaultMemberPermissions: &defaultAdmin,
	},
	CommandHandler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: r.ShootTurn(i.ChannelID, i.Member.User),
			},
		})
	},
}
