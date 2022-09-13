package repo

import (
	"context"
	"fmt"
	"time"

	"errors"

	d "github.com/holy-tech/discord-roulette/src/data"
	"go.mongodb.org/mongo-driver/bson"
)

func AcceptPlayer(channel string, user string) error {
	var result d.GameSettings
	accepted, err := GameIsAcceptedBy(channel, user)
	if err != nil {
		return err
	} else if accepted {
		return errors.New("you have already accepted")
	}

	db := Client.Database("games")
	gameCollection := db.Collection(fmt.Sprintf("%s_game", channel))
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()

	gameCollection.FindOne(ctx, bson.M{"channel": channel}).Decode(&result)
	update, _ := gameCollection.UpdateOne(ctx, bson.M{"channel": channel}, result)
	fmt.Println(result)
	fmt.Println(update)
	return nil
}
