package roulette

import (
	u "github.com/holy-tech/discord-roulette/src"
	db "github.com/holy-tech/discord-roulette/src/repo"
	_ "github.com/lib/pq"
)

func InitializeGame(channelId string) error {
	err := db.CreateGameDocument(channelId)
	u.CheckErr("Failed to create challenge", err)
	return nil
}
