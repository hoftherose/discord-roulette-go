package interfaces

//go:generate mockgen --destination=./mocks/table.go interfaces Game
type Game interface {
	TakeTurn()
	Accepted() bool
	GameFinished() bool
	Channel() string
}
