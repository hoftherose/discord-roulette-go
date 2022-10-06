package data

//go:generate mockgen --destination=./../../mocks/mentioner.go github.com/holy-tech/discord-roulette/src/data Mentioner
type Mentioner interface {
	Mention() string
}

//go:generate mockgen --destination=./../../mocks/user.go github.com/holy-tech/discord-roulette/src/data User
type User interface {
	GetID() string
	Mention() string
	Accept()
	HasAccepted() bool
}

type Player struct {
	Id       string `json:"id"`
	Accepted bool   `json:"accepted"`
}

func (p *Player) Mention() string {
	return "<@" + p.Id + ">"
}

func (p *Player) GetID() string {
	return p.Id
}

func (p *Player) Accept() {
	p.Accepted = true
}

func (p *Player) HasAccepted() bool {
	return p.Accepted
}
