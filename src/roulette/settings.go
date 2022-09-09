package roulette

import (
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	DefaultOpponent          *discordgo.User
	DefaultNumChamber        int64  = 6
	DefaultNumBullet         int64  = 1
	DefaultSpinChamber       bool   = false
	DefaultSpinChamberOnShot bool   = false
	DefaultReplaceBullet     bool   = false
	DefaultChannel           string = ""
)

type GameSettings struct {
	Opponent          *discordgo.User
	NumChamber        int64  `default:"6"`
	NumBullet         int64  `default:"1"`
	SpinChamber       bool   `default:"false"`
	SpinChamberOnShot bool   `default:"false"`
	ReplaceBullet     bool   `default:"false"`
	Channel           string `default:"none"`
}

var DefaultGameSettings GameSettings = GameSettings{
	DefaultOpponent,
	DefaultNumChamber,
	DefaultNumBullet,
	DefaultSpinChamber,
	DefaultSpinChamberOnShot,
	DefaultReplaceBullet,
	DefaultChannel,
}

func SettingToDoc(p *GameSettings) primitive.D {
	return bson.D{{"Players", 1}}
}
