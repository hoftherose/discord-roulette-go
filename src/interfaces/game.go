package interfaces

//go:generate mockgen --destination=./../../mocks/game.go github.com/holy-tech/discord-roulette/src/interfaces Game
type Game interface {
	StartGame()
	TakeTurn()
	IsAccepted() bool
	GameFinished() bool
	GetChannel() string
}
