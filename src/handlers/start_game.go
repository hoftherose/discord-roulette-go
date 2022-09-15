package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	d "github.com/holy-tech/discord-roulette/src/data"
	r "github.com/holy-tech/discord-roulette/src/roulette"
)

var defaultAdmin int64 = discordgo.PermissionAdministrator

func getSettingsFromOptions(
	s *discordgo.Session,
	opt []*discordgo.ApplicationCommandInteractionDataOption,
	challenger *discordgo.User,
	channel string,
) d.GameSettings {
	settings := d.GameSettings(d.DefaultGameSettings)

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(opt))
	for _, opt := range opt {
		optionMap[opt.Name] = opt
	}
	{
		if opponent, ok := optionMap["opponent"]; ok {
			settings.Opponents[challenger.ID] = d.Player{User: *challenger, Accepted: ""}
			settings.Opponents[opponent.UserValue(s).ID] = d.Player{User: *opponent.UserValue(s), Accepted: ""}
		}
		if numChamberValue, ok := optionMap["num_chambers"]; ok {
			settings.GunState.NumChamber = int(numChamberValue.IntValue())
		}
		if numBulletValue, ok := optionMap["num_bullets"]; ok {
			settings.GunState.NumBullets = int(numBulletValue.IntValue())
		}
		if spinChamberValue, ok := optionMap["spin_chamber"]; ok {
			settings.SpinChamberRule = spinChamberValue.BoolValue()
		}
		if spinChamberOnShotValue, ok := optionMap["spin_chamber_on_shot"]; ok {
			settings.SpinChamberOnShotRule = spinChamberOnShotValue.BoolValue()
		}
		if replaceBulletValue, ok := optionMap["replace_bullets"]; ok {
			settings.ReplaceBulletRule = replaceBulletValue.BoolValue()
		}
		settings.Channel = channel
	}
	return settings
}

var RouletteHandle = d.Handler{
	CommandSpecs: &discordgo.ApplicationCommand{
		Name:                     "roulette-start",
		Description:              "Roulette start game",
		DefaultMemberPermissions: &defaultAdmin,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "opponent",
				Description: fmt.Sprintf("Number of chambers in gun, defaults to %d", d.DefaultNumChamber),
				Type:        discordgo.ApplicationCommandOptionUser,
				Required:    true,
			},
			{
				Name:        "num_chambers",
				Description: fmt.Sprintf("Number of chambers in gun, defaults to %d", d.DefaultNumChamber),
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    false,
			},
			{
				Name:        "num_bullets",
				Description: fmt.Sprintf("Number of bullets in gun, defaults to %d", d.DefaultNumBullet),
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    false,
			},
			{
				Name:        "spin_chamber",
				Description: fmt.Sprintf("Spin chamber after pulling trigger, defaults to %t", d.DefaultSpinChamberRule),
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Required:    false,
			},
			{
				Name:        "spin_chamber_on_shot",
				Description: fmt.Sprintf("Spin chamber after gun fires, defaults to %t", d.DefaultSpinChamberOnShotRule),
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Required:    false,
			},
			{
				Name:        "replace_bullets",
				Description: fmt.Sprintf("Replace bullets if gun fires, defaults to %t", d.DefaultReplaceBulletRule),
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Required:    false,
			},
		},
	},
	CommandHandler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		challenger := i.Member.User
		settings := getSettingsFromOptions(s, options, challenger, i.ChannelID)

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: r.GameStart(&settings),
			},
		})
	},
}
