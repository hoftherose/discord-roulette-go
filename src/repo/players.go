package repo

import (
	"encoding/json"
	"fmt"

	"errors"

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

	result := GetGameDocument(channel)
	result.Opponents[0].Accepted = "True"
	err = UpdateGameDocument(bson.M{"channel": channel}, result, channel)
	if err != nil {
		return err
	}
	temp, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println(string(temp))
	return nil
}
