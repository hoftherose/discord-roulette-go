package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/holy-tech/discord-roulette/src/data"
	r "github.com/holy-tech/discord-roulette/src/roulette"
)

var GetGameHandle = data.Handler{
	CommandSpecs: &discordgo.ApplicationCommand{
		Name:                     "roulette-info",
		Description:              "Get roulette info",
		DefaultMemberPermissions: &defaultAdmin,
	},
	CommandHandler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		channel := i.ChannelID
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: r.GetGameInfo(channel),
			},
		})
	},
}
