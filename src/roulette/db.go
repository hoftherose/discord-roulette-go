package roulette

import (
	"fmt"
	"os"

	u "github.com/holy-tech/discord-roulette/src"
	db "github.com/holy-tech/discord-roulette/src/repo"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	db.CreateTable()
}

func InitializeGame(channelId string) string {
	challengeFile := fmt.Sprintf("%s_challenge.txt", channelId)
	_, err := os.Stat(challengeFile)
	u.CheckErr("Failed to create challenge", err)
	return "Challenge created in " + challengeFile
}
