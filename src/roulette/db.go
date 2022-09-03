package roulette

import (
	"fmt"
	"os"

	u "github.com/holy-tech/discord-roulette/src"
	_ "github.com/mattn/go-sqlite3"
)

func InitializeGame(channelId string) string {
	challengeFile := fmt.Sprintf("%s_challenge.txt", channelId)
	_, err := os.Stat(challengeFile)
	u.CheckErr("Failed to create challenge", err)
	return "Challenge created in " + challengeFile
}

// func RegisterStartGame(channelId string) {
// 	db.Connection
// }
