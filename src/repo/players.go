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
	update := UpdateGameDocument(bson.M{"channel": channel}, result, channel)
	temp, _ := json.MarshalIndent(result, "", "    ")
	temp2, _ := json.MarshalIndent(update, "", "    ")
	fmt.Println(string(temp))
	fmt.Println(string(temp2))
	return nil
}
