package roulette

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"github.com/holy-tech/discord-roulette/src/data"
	db "github.com/holy-tech/discord-roulette/src/repo"
)

func Shoot(s data.GameSettings, user *discordgo.User) (bool, error) {
	// GetCurrentPlayer(s)
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
	died, _ := Shoot(settings, user)
	if died {
		return "You died <@" + user.ID + ">"
	}
	return "You live <@" + user.ID + ">"
}
