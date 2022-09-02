package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var defaultAdmin int64 = discordgo.PermissionAdministrator
var (
	numChamber int64 = 6
	numBullet  int64 = 1
)

type gameSettings struct {
	numChamber int64
	numBullet  int64
}

func setDefaults(options []*discordgo.ApplicationCommandInteractionDataOption) gameSettings {

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}
	if numChamberValue, ok := optionMap["num_chambers"]; ok {
		numChamber = numChamberValue.IntValue()
	}
	if numBulletValue, ok := optionMap["num_chambers"]; ok {
		numBullet = numBulletValue.IntValue()
	}
	return gameSettings{numChamber, numBullet}
}

var RouletteHandle = Handler{
	&discordgo.ApplicationCommand{
		Name:                     "roulette-command",
		Description:              "Roulette command",
		DefaultMemberPermissions: &defaultAdmin,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "num_chambers",
				Description: "Number of chambers in gun, defaults to 6",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    false,
			},
			{
				Name:        "num_bullets",
				Description: "Number of bullets in gun, defaults to 1",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    false,
			},
		},
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		settings := setDefaults(options)

		resp := fmt.Sprintf("Preparing a %d-shooter with %d bullet(s).", settings.numChamber, settings.numBullet)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: resp,
			},
		})
	},
}
