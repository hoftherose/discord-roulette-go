package interfaces

//go:generate mockgen --destination=./mocks/table.go interfaces Game
type Game interface {
	StartGame()
	TakeTurn()
	IsAccepted() bool
	GameFinished() bool
	GetChannel() string
}
