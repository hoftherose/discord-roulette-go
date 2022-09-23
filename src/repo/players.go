package repo

import (
	"errors"
	"log"

	"github.com/bwmarrin/discordgo"
	u "github.com/holy-tech/discord-roulette/src"
)

func AcceptPlayer(channel string, user *discordgo.User) error {
	accepted, err := GameIsAcceptedBy(channel, user)
	if err != nil {
		return err
	} else if accepted {
		return errors.New("you have already accepted")
	}

	result, _ := GetGameDocument(channel)

	for k, opponent := range result.Opponents {
		if k == user.ID {
			opponent.Accepted = "true"
			result.Opponents[k] = opponent
			break
		}
	}
	err = UpdateGameDocument(channel, result)
	if err != nil {
		return err
	}
	return nil
}

func AwaitingPlayer(channel string) (string, bool) {
	var awaitingPlayers []string
	result, err := GetGameDocument(channel)
	if err != nil {
		log.Fatalf("Could not get game: %v", err)
	}
	for i, opponent := range result.Opponents {
		if opponent.Accepted != "true" {
			awaitingPlayers = append(awaitingPlayers, result.Opponents[i].ID)
		}
	}
	if len(awaitingPlayers) == 0 {
		err := AcceptGame(channel)
		if err != nil {
			return err.Error(), false
		}
		return "All players have accepted!", true
	}
	return "Still waiting for <@" + u.JoinStrings(">, <@", awaitingPlayers...) + ">", false
}

func GetCurrentPlayer(channel string) string {
	settings, err := GetGameDocument(channel)
	if err != nil {
		log.Fatalf("Could not get game: %v", err)
	}
	return settings.TableState.GetCurrentPlayer()
}
