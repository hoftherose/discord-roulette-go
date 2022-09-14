package repo

import (
	"errors"
	"log"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

func AcceptPlayer(channel string, user *discordgo.User) error {
	accepted, err := GameIsAcceptedBy(channel, user)
	if err != nil {
		return err
	} else if accepted {
		return errors.New("you have already accepted")
	}

	result, err := GetGameDocument(channel)
	if err != nil {
		log.Fatalf("Could not get game: %v", err)
	}

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
