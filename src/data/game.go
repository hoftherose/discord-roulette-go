package data

import (
	i "github.com/holy-tech/discord-roulette/src/interfaces"
)

var (
	DefaultGameAccepted bool   = false
	DefaultChannel      string = ""
)

type GameStatus struct {
	Table        i.Table
	Revolver     i.Gun
	GameAccepted bool   `json:"game_accepted,omitempty"`
	Channel      string `json:"channel,omitempty"`
}

var DefaultGameStatus GameStatus = GameStatus{
	DefaultGameTable,
	DefaultRevolver,
	DefaultGameAccepted,
	DefaultChannel,
}

func (s *GameStatus) StartGame() {
	//TODO implement
}

func (s *GameStatus) TakeTurn() {
	//TODO implement
}

func (s *GameStatus) IsAccepted() bool {
	for _, player := range s.Table.GetSeating() {
		if !player.HasAccepted() {
			return false
		}
	}
	return true
}

func (s *GameStatus) GameFinished() bool {
	return s.Table.NumPlayers() < 2
}

func (s *GameStatus) GetChannel() string {
	return s.Channel
}
