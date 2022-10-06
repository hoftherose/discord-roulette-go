package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	data "github.com/holy-tech/discord-roulette/src/data"
	r "github.com/holy-tech/discord-roulette/src/roulette"
)

var defaultAdmin int64 = discordgo.PermissionAdministrator

func getGameFromOptions(
	s *discordgo.Session,
	opt []*discordgo.ApplicationCommandInteractionDataOption,
	challenger *discordgo.User,
	channel string,
) data.GameStatus {
	var numChamber int = data.DefaultNumChamber
	var numBullets int = data.DefaultNumBullet
	var players []data.User = []data.User{}

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(opt))
	for _, opt := range opt {
		optionMap[opt.Name] = opt
	}

	if opponent, ok := optionMap["opponent"]; ok {
		opponent := opponent.UserValue(s)
		players = append(players, &data.Player{Id: opponent.ID, Accepted: data.DefaultGameAccepted})
	}
	players = append(players, &data.Player{Id: challenger.ID, Accepted: data.DefaultGameAccepted})
	if numChamberValue, ok := optionMap["num_chambers"]; ok {
		numChamber = int(numChamberValue.IntValue())
	}
	if numBulletValue, ok := optionMap["num_bullets"]; ok {
		numBullets = int(numBulletValue.IntValue())
	}
	table := data.GameTable{}
	table.InitTable(players...)
	revolver := data.Revolver{}
	revolver.ReloadGun(numChamber, numBullets)
	game := data.GameStatus{
		Table:        &table,
		Revolver:     &revolver,
		GameAccepted: data.DefaultGameAccepted,
		Channel:      channel,
	}
	return game
}

var RouletteHandle = Handler{
	CommandSpecs: &discordgo.ApplicationCommand{
		Name:                     "roulette-start",
		Description:              "Roulette start game",
		DefaultMemberPermissions: &defaultAdmin,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "opponent",
				Description: fmt.Sprintf("Number of chambers in gun, defaults to %d", data.DefaultNumChamber),
				Type:        discordgo.ApplicationCommandOptionUser,
				Required:    true,
			},
			{
				Name:        "num_chambers",
				Description: fmt.Sprintf("Number of chambers in gun, defaults to %d", data.DefaultNumChamber),
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    false,
			},
			{
				Name:        "num_bullets",
				Description: fmt.Sprintf("Number of bullets in gun, defaults to %d", data.DefaultNumBullet),
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    false,
			},
		},
	},
	CommandHandler: func(s *discordgo.Session, ic *discordgo.InteractionCreate) {
		options := ic.ApplicationCommandData().Options
		challenger := ic.Member.User
		game := getGameFromOptions(s, options, challenger, ic.ChannelID)

		s.InteractionRespond(ic.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: r.GameStart(&game),
			},
		})
	},
}
