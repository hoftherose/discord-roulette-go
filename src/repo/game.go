package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	u "github.com/holy-tech/discord-roulette/src"
	d "github.com/holy-tech/discord-roulette/src/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateGameDocument(channel string, settings interface{}) error {
	var result bson.M
	db := Client.Database("games")
	gameCollection := db.Collection(fmt.Sprintf("%s_game", channel))
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()

	gameCollection.FindOne(ctx, bson.D{}).Decode(&result)

	if result == nil {
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

func GetGameDocument(channel string) d.GameSettings {
	var result d.GameSettings
	db := Client.Database("games")
	gameCollection := db.Collection(fmt.Sprintf("%s_game", channel))
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()

	gameCollection.FindOne(ctx, bson.M{"channel": channel}).Decode(&result)
	return result
}

func UpdateGameDocument(filter interface{}, update interface{}, channel string) *mongo.UpdateResult {
	db := Client.Database("games")
	gameCollection := db.Collection(fmt.Sprintf("%s_game", channel))
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()

	updated, _ := gameCollection.UpdateOne(ctx, bson.M{"channel": channel}, update)
	return updated
}

func GameIsAcceptedBy(channel string, user string) (bool, error) {
	var result bson.M
	db := Client.Database("games")
	gameCollection := db.Collection(fmt.Sprintf("%s_game", channel))
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()

	gameCollection.FindOne(ctx, bson.M{"channel": channel}).Decode(&result)
	// TODO look for specific player
	if len(result) == 0 {
		return false, errors.New("game not found")
	}
	temp := result["opponents"].(bson.A)[0].(bson.M)["accepted"]
	return temp != "", nil
}
