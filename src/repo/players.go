package repo

import (
	"fmt"

	"errors"

	"github.com/bwmarrin/discordgo"
	d "github.com/holy-tech/discord-roulette/src/data"
	"go.mongodb.org/mongo-driver/bson"
)

func AcceptPlayer(channel string, user *discordgo.User) error {
	var result d.GameSettings
	accepted, err := GameIsAcceptedBy(channel, user)
	if err != nil {
		return err
	} else if accepted {
		return errors.New("you have already accepted")
	}

	update := UpdateGameDocument(bson.M{"channel": channel}, result, channel)
	fmt.Println(result)
	fmt.Println(update)
	return nil
}
