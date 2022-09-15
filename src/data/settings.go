package data

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

var (
	DefaultOpponents         map[string]Player = map[string]Player{}
	DefaultGameAccepted      bool              = false
	DefaultSpinChamber       bool              = false
	DefaultSpinChamberOnShot bool              = false
	DefaultReplaceBullet     bool              = false
	DefaultChannel           string            = ""
)

type Player struct {
	discordgo.User
	Accepted string `json:"accepted"`
}

type GameSettings struct {
	Opponents         map[string]Player `json:"opponent,omitempty"`
	TableState        TableState        `json:"num_chambers,omitempty"`
	GunState          GunState          `json:"num_bullets,omitempty"`
	GameAccepted      bool              `json:"game_accepted,omitempty"`
	SpinChamber       bool              `json:"spin_chamber,omitempty"`
	SpinChamberOnShot bool              `json:"spin_chamber_on_shot,omitempty"`
	ReplaceBullet     bool              `json:"replace_bullet,omitempty"`
	Channel           string            `json:"channel,omitempty"`
}

var DefaultGameSettings GameSettings = GameSettings{
	DefaultOpponents,
	DefaultTableState,
	DefaultGunState,
	DefaultGameAccepted,
	DefaultSpinChamber,
	DefaultSpinChamberOnShot,
	DefaultReplaceBullet,
	DefaultChannel,
}

func (s *GameSettings) GetCurrentPlayer() string {
	curr_player := s.TableState.CurrentTurn
	return s.TableState.Turns[curr_player]
}

func (s *GameSettings) SetNextPlayer() {
	s.TableState.CurrentTurn = s.TableState.CurrentTurn % int64(len(s.TableState.Turns))
}

func (s *GameSettings) Shoot(user *discordgo.User) (bool, error) {
	curr_player := s.GetCurrentPlayer()
	if user.ID != curr_player {
		return false, errors.New("it is not your turn")
	}
	curr_chamber := s.GunState.CurrentChamber
	died := s.GunState.Chambers[curr_chamber]
	s.GunState.CurrentChamber = (curr_chamber + 1) % int64(s.GunState.NumChamber)
	if died {
		// TODO Setup actual loser table
		s.TableState.Losers = append(s.TableState.Losers, true)
	}
	return died, nil
}
