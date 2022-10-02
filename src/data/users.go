package data

import (
	"github.com/bwmarrin/discordgo"
)

type Player struct {
	discordgo.User
	Accepted bool `json:"accepted"`
}

func (p *Player) GetID() string {
	return p.ID
}

func (p *Player) Accept() {
	p.Accepted = true
}

func (p *Player) HasAccepted() bool {
	return p.Accepted
}
