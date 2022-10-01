package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"log"

	d "github.com/holy-tech/discord-roulette/src/data"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateGameDocument(channel string, settings interface{}) error {
	var result bson.M
	db := Client.Database("games")
	gameCollection := db.Collection(fmt.Sprintf("%s_game", channel))
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()

	gameCollection.FindOne(ctx, bson.D{}).Decode(&result)

	if result != nil {
		return errors.New("game already exists")
	}
	_, err := gameCollection.InsertOne(ctx, settings)
	if err != nil {
		log.Fatalf("Error executing query: %v\n", err)
	}
	return err
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

func GetGameDocument(channel string) (d.GameSettings, error) {
	var result d.GameSettings
	db := Client.Database("games")
	gameCollection := db.Collection(fmt.Sprintf("%s_game", channel))
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()

	encResult := gameCollection.FindOne(ctx, bson.M{"channel": channel})
	encResult.Decode(&result)
	return result, encResult.Err()
}

func UpdateGameDocument(channel string, new interface{}) error {
	db := Client.Database("games")
	gameCollection := db.Collection(fmt.Sprintf("%s_game", channel))
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()

	update, err := gameCollection.ReplaceOne(ctx, bson.M{"channel": channel}, new)

	if err != nil {
		return err
	} else if update.ModifiedCount == 0 {
		return errors.New("no update occured")
	}
	return nil
}
