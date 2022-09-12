package roulette

import (
	db "github.com/holy-tech/discord-roulette/src/repo"
)

func Died() bool {
	return true
}

func ShootTurn(channel string, user string) string {
	if !db.GameIsAcceptedBy(channel, user) {
		return "Game still is not accepted"
	}
	if Died() {
		return "You died <@" + user + ">"
	}
	return "You live <@" + user + ">"
}
