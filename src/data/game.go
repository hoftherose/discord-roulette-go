package data

import (
	"github.com/bwmarrin/discordgo"
	i "github.com/holy-tech/discord-roulette/src/interfaces"
)

var (
	DefaultGameAccepted bool   = false
	DefaultChannel      string = ""
)

type Player struct {
	discordgo.User
	Accepted bool `json:"accepted"`
}

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

func (s *GameStatus) TakeTurn(user *discordgo.User) (bool, error) {
	// currPlayer := s.TableState.GetCurrentPlayer()
	// if user.ID != currPlayer {
	// 	return false, errors.New("it is not your turn")
	// }
	// shot := s.Revolver.Chambers[s.Revolver.CurrentChamber]
	// s.Revolver.SetNextChamber()
	// s.TableState.SetNextPlayer()
	// if shot {
	// 	delete(s.Opponents, user.ID)
	// 	s.TableState.RemovePlayer(user.ID)
	// }
	// s.Revolver.ClearChamber(shot)
	// if s.Revolver.NumBulletsLeft <= 0 {
	// 	s.Revolver.SpinChamber()
	// }
	return false, nil
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

func (p *Player) GetID() string {
	return p.ID
}

func (p *Player) Accept() {
	p.Accepted = true
}

func (p *Player) HasAccepted() bool {
	return p.Accepted
}
