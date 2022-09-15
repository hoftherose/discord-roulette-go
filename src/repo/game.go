package repo

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

func AcceptGame(channel string) error {
	result, _ := GetGameDocument(channel)
	result.GameAccepted = true
	err := UpdateGameDocument(channel, result)
	return err
}

func GameIsAccepted(channel string) (bool, error) {
	result, err := GetGameDocument(channel)

	if err != nil {
		return false, errors.New("could not find game")
	}

	return result.GameAccepted, nil
}

func GameIsAcceptedBy(channel string, user *discordgo.User) (bool, error) {
	result, err := GetGameDocument(channel)

	if err != nil {
		return false, errors.New("could not find game")
	}

	for i, opponent := range result.Opponents {
		if opponent.ID == user.ID {
			return result.Opponents[i].Accepted == "true", nil
		}
	}
	return false, errors.New("could not find user in game")
}
