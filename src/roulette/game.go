package roulette

import (
	"github.com/bwmarrin/discordgo"
	"github.com/holy-tech/discord-roulette/src/data"
	db "github.com/holy-tech/discord-roulette/src/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetGun(revolver primitive.M) data.Gun {
	chamber := []bool{}
	for _, bullet := range revolver["chamber"].(primitive.A) {
		chamber = append(chamber, bullet.(bool))
	}
	numBulletsLeft := int(revolver["numbulletsleft"].(int32))
	currentChamber := int(revolver["currentchamber"].(int32))
	seed := revolver["seed"].(int64)
	gun := &data.Revolver{
		Chamber:        chamber,
		NumBulletsLeft: numBulletsLeft,
		CurrentChamber: currentChamber,
		Seed:           seed,
	}
	return gun
}

func GetTable(table primitive.M) data.Table {
	seating := []data.User{}
	for _, player := range table["seating"].(primitive.A) {
		user := GetUser(player.(primitive.M))
		seating = append(seating, user)
	}
	currentTurn := int(table["currentturn"].(int32))
	seed := table["seed"].(int64)
	gameTable := &data.GameTable{
		Seating:     seating,
		CurrentTurn: currentTurn,
		Seed:        seed,
	}
	return gameTable
}

func GetUser(user primitive.M) data.User {
	id := user["id"].(string)
	accepted := user["accepted"].(bool)
	player := &data.Player{
		Id:       id,
		Accepted: accepted,
	}
	return player
}

func GetGame(channel string) (data.GameStatus, error) {
	var result bson.M
	encResult := db.GetGameDocument(channel)

	encResult.Decode(&result)
	gun := GetGun(result["revolver"].(bson.M))
	table := GetTable(result["table"].(bson.M))
	game := data.GameStatus{
		Table:        table,
		Revolver:     gun,
		GameAccepted: result["gameaccepted"].(bool),
		Channel:      result["channel"].(string),
	}
	return game, encResult.Err()
}

func GetGameInfo(channel string) string {
	result, _ := GetGame(channel)
	return "Info" + result.Channel
}

func ShootTurn(channel string, user *discordgo.User) string {
	var message string
	game, err := GetGame(channel)
	if err != nil {
		return "No game found"
	}
	accepted := game.IsAccepted()
	if !accepted {
		return "Game still is not accepted"
	}
	if game.Table.GetCurrentPlayer().GetID() != user.ID {
		return "It is not your turn"
	}

	shot := game.TakeTurn()
	db.UpdateGameDocument(channel, game)
	if err != nil {
		message = "Error: " + err.Error()
	} else if shot {
		message = "You died " + user.Mention()
	} else {
		message = "You live " + user.Mention()
	}
	if !game.GameFinished() {
		return message + "\nIt is " + game.Table.GetCurrentPlayer().Mention() + " turn."
	}
	db.DeleteGameDocument(channel)
	return message + "\nThe winner is: " + game.Table.GetCurrentPlayer().Mention()
}
