package data

import (
	"github.com/bwmarrin/discordgo"
)

var (
	DefaultOpponents         []Player
	DefaultNumChamber        int64  = 6
	DefaultNumBullet         int64  = 1
	DefaultSpinChamber       bool   = false
	DefaultSpinChamberOnShot bool   = false
	DefaultReplaceBullet     bool   = false
	DefaultChannel           string = ""
)

type Player struct {
	discordgo.User
	Accepted string `json:"accepted"`
}

type GameSettings struct {
	Opponents         []Player `json:"opponent,omitempty"`
	NumChamber        int64    `json:"num_chambers,omitempty"`
	NumBullet         int64    `json:"num_bullets,omitempty"`
	SpinChamber       bool     `json:"spin_chamber,omitempty"`
	SpinChamberOnShot bool     `json:"spin_chamber_on_shot,omitempty"`
	ReplaceBullet     bool     `json:"replace_bullet,omitempty"`
	Channel           string   `json:"channel,omitempty"`
}

var DefaultGameSettings GameSettings = GameSettings{
	DefaultOpponents,
	DefaultNumChamber,
	DefaultNumBullet,
	DefaultSpinChamber,
	DefaultSpinChamberOnShot,
	DefaultReplaceBullet,
	DefaultChannel,
}