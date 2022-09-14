package repo

import (
	"errors"

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

	result, err := GetGameDocument(channel)
	u.CheckErr("Could not get game: %v", err)

	result.Opponents[0].Accepted = "True"
	err = UpdateGameDocument(bson.M{"channel": channel}, result, channel)
	if err != nil {
		return err
	}
	return nil
}
