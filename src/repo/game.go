package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	u "github.com/holy-tech/discord-roulette/src"
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

func UpdateGameDocument(filter interface{}, new interface{}, channel string) error {
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

func GameIsAcceptedBy(channel string, user *discordgo.User) (bool, error) {
	result := GetGameDocument(channel)
	fmt.Println(result)

	// TODO look for specific player
	// if len(result) == 0 {
	// 	return false, errors.New("game not found")
	// }
	temp := result.Opponents[0].Accepted
	fmt.Println("temp")
	fmt.Println(temp)
	return temp != "", nil
}
