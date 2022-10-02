package interfaces

//go:generate mockgen --destination=./../../mocks/user.go github.com/holy-tech/discord-roulette/src/interfaces User
type User interface {
	GetID() string
	Mention() string
	Accept()
	HasAccepted() bool
}
