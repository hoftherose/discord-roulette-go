package roulette

func Died() bool {
	return true
}

func ShootTurn(channel string, user string) string {
	if Died() {
		return "You died <@" + user + ">"
	}
	return "You live <@" + user + ">"
}
