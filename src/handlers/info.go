package handlers

import (
	"github.com/bwmarrin/discordgo"
	embed "github.com/holy-tech/discord-roulette/src/embed"
	i "github.com/holy-tech/discord-roulette/src/interfaces"
	r "github.com/holy-tech/discord-roulette/src/roulette"
)

var GetGameHandle = i.Handler{
	CommandSpecs: &discordgo.ApplicationCommand{
		Name:                     "roulette-info",
		Description:              "Get roulette info",
		DefaultMemberPermissions: &defaultAdmin,
	},
	CommandHandler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		channel := i.ChannelID
		embed.TempEmbed(s, channel)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: r.GetGameInfo(channel),
			},
		})
	},
}
