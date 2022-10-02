package interfaces

//go:generate mockgen --destination=./mocks/player.go interfaces Player
type User interface {
	GetID() string
	Mention() string
	Accept()
	Accepted() bool
}
