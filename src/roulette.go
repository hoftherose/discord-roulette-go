package handlers

import "github.com/bwmarrin/discordgo"

var RouletteHandle = Handler{
	&discordgo.ApplicationCommand{
		Name:        "roulette-command",
		Description: "Roulette command",
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hey there! Congratulations, you just executed your first slash command",
			},
		})
	},
}
