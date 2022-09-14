package repo

import (
	"errors"
	"log"

	"github.com/bwmarrin/discordgo"
	u "github.com/holy-tech/discord-roulette/src"
	"go.mongodb.org/mongo-driver/bson"
)

func AcceptPlayer(channel string, user *discordgo.User) error {
	accepted, err := GameIsAcceptedBy(channel, user)
	if err != nil {
		return err
	} else if accepted {
		return errors.New("you have already accepted")
	}

	result, _ := GetGameDocument(channel)

	for i, opponent := range result.Opponents {
		if opponent.ID == user.ID {
			result.Opponents[i].Accepted = "true"
			break
		}
	}
	err = UpdateGameDocument(bson.M{"channel": channel}, result, channel)
	if err != nil {
		return err
	}
	return nil
}

func AwaitingPlayer(channel string) string {
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
			return err.Error()
		}
		return "All players have accepted!"
	}
	return "Still waiting for <@" + u.JoinStrings(">, <@", awaitingPlayers...) + ">"
}
