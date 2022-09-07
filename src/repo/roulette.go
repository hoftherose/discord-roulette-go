package repo

import (
	"context"
	"fmt"
	"time"

	u "github.com/holy-tech/discord-roulette/src"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateTable() {
	db := Client.Database("games")
	gameCollection := db.Collection("channel_name_game")
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()
	result, err := gameCollection.InsertOne(ctx, bson.D{{"x", 1}})
	u.CheckErr("Error executing query: %v", err)
	fmt.Printf("Table created successfully: %v", result)
}
