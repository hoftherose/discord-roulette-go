package repo

import (
	"errors"
)

func AcceptPlayer(channel string, user string) error {
	if GameIsAcceptedBy(channel, user) {
		return errors.New("you have already accepted")
	}

	// db := Client.Database("games")
	// gameCollection := db.Collection(fmt.Sprintf("%s_game", channel))
	// ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancelCtx()

	// result, _ := gameCollection.UpdateOne(ctx, bson.M{"channel": channel}, bson.D{{"opponents": [{"accepted": "true"}]}})
	// fmt.Println(result)
	return nil
}
