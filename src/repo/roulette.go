package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	u "github.com/holy-tech/discord-roulette/src"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateGameDocument(channel string) error {
	var result bson.M
	db := Client.Database("games")
	gameCollection := db.Collection(fmt.Sprintf("%s_game", channel))
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()

	gameCollection.FindOne(ctx, bson.D{{"x", 1}}).Decode(&result)
	fmt.Println(result)

	if result == nil {
		result, err := gameCollection.InsertOne(ctx, bson.D{{"x", 1}})
		u.CheckErr("Error executing query: %v\n", err)
		fmt.Printf("Table created successfully: %v\n", result)
		return nil
	}
	return errors.New("Game already exists")
}
