package data

import (
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
