package interfaces

//go:generate mockgen --destination=./mocks/table.go interfaces Game
type Game interface {
	TakeTurn()
	Accepted()
	GameFinished() bool
}
