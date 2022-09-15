package roulette

import (
	"math/rand"
	"time"

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
	var k int64
	s.GunState.Chambers = make([]bool, s.GunState.NumChamber)
	for k = 0; k < s.GunState.NumChamber; k++ {
		s.GunState.Chambers[k] = k < s.GunState.NumBullets
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s.GunState.Chambers), func(i, j int) {
		s.GunState.Chambers[i], s.GunState.Chambers[j] = s.GunState.Chambers[j], s.GunState.Chambers[i]
	})
	db.UpdateGameDocument(channel, s)
}
