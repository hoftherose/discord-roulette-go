package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	u "github.com/holy-tech/discord-roulette/src"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateGameDocument(channel string, settings primitive.D) error {
	var result bson.M
	db := Client.Database("games")
	gameCollection := db.Collection(fmt.Sprintf("%s_game", channel))
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()

	gameCollection.FindOne(ctx, bson.D{}).Decode(&result)

	if result == nil {
		settings := bson.D{{"Players", 1}}
		result, err := gameCollection.InsertOne(ctx, settings)
		u.CheckErr("Error executing query: %v\n", err)
		fmt.Printf("Table created successfully: %v\n", result)
		return nil
	}
	return errors.New("game already exists")
}

func DeleteGameDocument(channel string) error {
	db := Client.Database("games")
	gameCollection := db.Collection(fmt.Sprintf("%s_game", channel))
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()

	result, err := gameCollection.DeleteOne(ctx, bson.D{{}})
	if result.DeletedCount != 1 {
		return errors.New("no game is currently ongoing")
	}
	return err

}
