package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	r "github.com/holy-tech/discord-roulette/src/roulette"
)

var defaultAdmin int64 = discordgo.PermissionAdministrator

func getSettingsFromOptions(
	s *discordgo.Session,
	opt []*discordgo.ApplicationCommandInteractionDataOption,
	challenger *discordgo.User,
	channel string,
) r.GameSettings {
	settings := r.GameSettings(r.DefaultGameSettings)

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(opt))
	for _, opt := range opt {
		optionMap[opt.Name] = opt
	}
	{
		if opponent, ok := optionMap["opponent"]; ok {
			settings.Opponents = append(settings.Opponents, r.Player{User: *challenger, Accepted: ""})
			settings.Opponents = append(settings.Opponents, r.Player{User: *opponent.UserValue(s), Accepted: ""})
		}
		if numChamberValue, ok := optionMap["num_chambers"]; ok {
			settings.NumChamber = numChamberValue.IntValue()
		}
		if numBulletValue, ok := optionMap["num_bullets"]; ok {
			settings.NumBullet = numBulletValue.IntValue()
		}
		if spinChamberValue, ok := optionMap["spin_chamber"]; ok {
			settings.SpinChamber = spinChamberValue.BoolValue()
		}
		if spinChamberOnShotValue, ok := optionMap["spin_chamber_on_shot"]; ok {
			settings.SpinChamberOnShot = spinChamberOnShotValue.BoolValue()
		}
		if replaceBulletValue, ok := optionMap["replace_bullets"]; ok {
			settings.ReplaceBullet = replaceBulletValue.BoolValue()
		}
		settings.Channel = channel
	}
	return settings
}

var RouletteHandle = Handler{
	&discordgo.ApplicationCommand{
		Name:                     "roulette-start",
		Description:              "Roulette start game",
		DefaultMemberPermissions: &defaultAdmin,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "opponent",
				Description: fmt.Sprintf("Number of chambers in gun, defaults to %d", r.DefaultNumChamber),
				Type:        discordgo.ApplicationCommandOptionUser,
				Required:    true,
			},
			{
				Name:        "num_chambers",
				Description: fmt.Sprintf("Number of chambers in gun, defaults to %d", r.DefaultNumChamber),
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    false,
			},
			{
				Name:        "num_bullets",
				Description: fmt.Sprintf("Number of bullets in gun, defaults to %d", r.DefaultNumBullet),
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    false,
			},
			{
				Name:        "spin_chamber",
				Description: fmt.Sprintf("Spin chamber after pulling trigger, defaults to %t", r.DefaultSpinChamber),
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Required:    false,
			},
			{
				Name:        "spin_chamber_on_shot",
				Description: fmt.Sprintf("Spin chamber after gun fires, defaults to %t", r.DefaultSpinChamberOnShot),
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Required:    false,
			},
			{
				Name:        "replace_bullets",
				Description: fmt.Sprintf("Replace bullets if gun fires, defaults to %t", r.DefaultReplaceBullet),
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Required:    false,
			},
		},
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
