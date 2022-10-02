package interfaces

//go:generate mockgen --destination=./mocks/player.go interfaces Player
type Player interface {
	GetID()
	Mention()
	Accepted()
}
