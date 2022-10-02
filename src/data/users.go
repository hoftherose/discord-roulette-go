package data

import (
	"github.com/bwmarrin/discordgo"
)

type User struct {
	discordgo.User
	accepted bool
}

func (u *User) GetID() string {
	return u.ID
}

func (u *User) Accept() {
	u.accepted = true
}

func (u *User) Accepted() bool {
	return u.accepted
}
