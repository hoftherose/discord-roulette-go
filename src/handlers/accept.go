package handlers

import (
	"github.com/bwmarrin/discordgo"
	r "github.com/holy-tech/discord-roulette/src/roulette"
)

var AcceptHandle = Handler{
	&discordgo.ApplicationCommand{
		Name:                     "roulette-accept",
		Description:              "Accept roulette match",
		DefaultMemberPermissions: &defaultAdmin,
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: r.ChallengeAccept(&r.GameSettings{Channel: i.ChannelID}),
			},
		})
	},
}

var DenyHandle = Handler{
	&discordgo.ApplicationCommand{
		Name:                     "roulette-deny",
		Description:              "Deny roulette match",
		DefaultMemberPermissions: &defaultAdmin,
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: r.ChallengeDeny(&r.GameSettings{Channel: i.ChannelID}),
			},
		})
	},
}
