package roulette

import (
	"errors"
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"github.com/holy-tech/discord-roulette/src/data"
	db "github.com/holy-tech/discord-roulette/src/repo"
)

func Shoot(s data.GameSettings, user *discordgo.User) (bool, error) {
	curr_player := s.GetCurrentPlayer()
	if user.ID != curr_player {
		return false, errors.New("it is not your turn")
	}
	died := rand.Intn(int(s.GunState.NumChamber)) > int(s.GunState.NumBullets)
	// if died {

	// 	s.TableState.Losers = append(s.TableState.Losers, )
	// }
	return died, nil
}

func ShootTurn(channel string, user *discordgo.User) string {
	accepted, err := db.GameIsAccepted(channel)
	if err != nil {
		return "No shots fired: " + err.Error()
	}
	settings, _ := db.GetGameDocument(channel)
	if !accepted {
		return "Game still is not accepted"
	}
	died, err := Shoot(settings, user)
	if err != nil {
		return "Error: " + err.Error()
	}
	if died {
		return "You died <@" + user.ID + ">"
	}
	return "You live <@" + user.ID + ">"
}
