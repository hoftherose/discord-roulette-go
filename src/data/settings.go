package data

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

var (
	DefaultOpponents             map[string]Player = map[string]Player{}
	DefaultGameAccepted          bool              = false
	DefaultSpinChamberRule       bool              = false
	DefaultSpinChamberOnShotRule bool              = false
	DefaultReplaceBulletRule     bool              = false
	DefaultChannel               string            = ""
)

type Player struct {
	discordgo.User
	Accepted string `json:"accepted"`
}

type GameSettings struct {
	Opponents             map[string]Player
	TableState            TableState
	GunState              GunState
	GameAccepted          bool   `json:"game_accepted,omitempty"`
	SpinChamberRule       bool   `json:"spin_chamber,omitempty"`
	SpinChamberOnShotRule bool   `json:"spin_chamber_on_shot,omitempty"`
	ReplaceBulletRule     bool   `json:"replace_bullet,omitempty"`
	Channel               string `json:"channel,omitempty"`
}

var DefaultGameSettings GameSettings = GameSettings{
	DefaultOpponents,
	DefaultTableState,
	DefaultGunState,
	DefaultGameAccepted,
	DefaultSpinChamberRule,
	DefaultSpinChamberOnShotRule,
	DefaultReplaceBulletRule,
	DefaultChannel,
}

func (s *GameSettings) Shoot(user *discordgo.User) (bool, error) {
	curr_player := s.TableState.GetCurrentPlayer()
	if user.ID != curr_player {
		return false, errors.New("it is not your turn")
	}
	died := s.GunState.Chambers[s.GunState.CurrentChamber]
	s.GunState.SetNextChamber()
	if died {
		// TODO Setup actual loser table
		s.TableState.Losers = append(s.TableState.Losers, true)
	}
	return died, nil
}
