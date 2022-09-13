package roulette

import (
	db "github.com/holy-tech/discord-roulette/src/repo"
)

func Died() bool {
	return true
}

func ShootTurn(channel string, user string) string {
	accepted, err := db.GameIsAcceptedBy(channel, user)
	if err != nil {
		return err.Error()
	}
	if !accepted {
		return "Game still is not accepted"
	} else if Died() {
		return "You died <@" + user + ">"
	}
	return "You live <@" + user + ">"
}
