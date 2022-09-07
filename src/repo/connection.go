package repo

import (
	"context"

	u "github.com/holy-tech/discord-roulette/src"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var uri = "mongodb://root:root@localhost:27017"

func init() {
	var err error
	Client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	u.CheckErr("Could not return DB connection: %v", err)
}
