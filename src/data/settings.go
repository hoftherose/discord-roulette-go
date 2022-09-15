package data

import (
	"errors"
	"math/rand"
	"time"

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
	Opponents             map[string]Player `json:"opponent,omitempty"`
	TableState            TableState        `json:"num_chambers,omitempty"`
	GunState              GunState          `json:"num_bullets,omitempty"`
	GameAccepted          bool              `json:"game_accepted,omitempty"`
	SpinChamberRule       bool              `json:"spin_chamber,omitempty"`
	SpinChamberOnShotRule bool              `json:"spin_chamber_on_shot,omitempty"`
	ReplaceBulletRule     bool              `json:"replace_bullet,omitempty"`
	Channel               string            `json:"channel,omitempty"`
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

func (s *GameSettings) GetCurrentPlayer() string {
	curr_player := s.TableState.CurrentTurn
	return s.TableState.Turns[curr_player]
}

func (s *GameSettings) SetNextPlayer() {
	s.TableState.CurrentTurn = s.TableState.CurrentTurn % len(s.TableState.Turns)
}

func (s *GameSettings) Shoot(user *discordgo.User) (bool, error) {
	curr_player := s.GetCurrentPlayer()
	if user.ID != curr_player {
		return false, errors.New("it is not your turn")
	}
	curr_chamber := s.GunState.CurrentChamber
	died := s.GunState.Chambers[curr_chamber]
	s.GunState.CurrentChamber = (curr_chamber + 1) % s.GunState.NumChamber
	if died {
		// TODO Setup actual loser table
		s.TableState.Losers = append(s.TableState.Losers, true)
	}
	return died, nil
}

func (s *GameSettings) SpinChamber() {
	s.GunState.Chambers = make([]bool, s.GunState.NumChamber)
	for k := 0; k < s.GunState.NumChamber; k++ {
		s.GunState.Chambers[k] = k < s.GunState.NumBullets
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s.GunState.Chambers), func(i, j int) {
		s.GunState.Chambers[i], s.GunState.Chambers[j] = s.GunState.Chambers[j], s.GunState.Chambers[i]
	})
}
