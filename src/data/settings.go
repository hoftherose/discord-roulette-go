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
	shot := s.GunState.Chambers[s.GunState.CurrentChamber]
	s.GunState.SetNextChamber()
	s.TableState.SetNextPlayer()
	if shot {
		s.TableState.Losers = append(s.TableState.Losers, user.ID)
		delete(s.Opponents, user.ID)
		s.GunState.NumBulletsLeft--
	}
	if s.GunState.NumBulletsLeft <= 0 {
		s.GunState.SpinChamber()
	}
	return shot, nil
}
