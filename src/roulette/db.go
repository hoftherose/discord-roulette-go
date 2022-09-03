package roulette

import (
	"fmt"
	"os"
	// sql "github.com/mattn/go-sqlite3"
)

func RegisterStartGame(channelId string) (string, error) {
	challengeFile := fmt.Sprintf("%s_challenge.txt", channelId)
	if _, err := os.Stat(challengeFile); err != nil {
		return "Failed to create challenge", err
	}
	return "Challenge created in " + challengeFile, nil
}
