package roulette

import (
	db "github.com/holy-tech/discord-roulette/src/repo"
)

func SetTable(channel string) {
	s, _ := db.GetGameDocument(channel)
	s.TableState.Turns = make([]string, len(s.Opponents))
	i := 0
	for k := range s.Opponents {
		s.TableState.Turns[i] = k
		i++
	}
	// TODO Set random seating
	s.SpinChamber()
	db.UpdateGameDocument(channel, s)
}
