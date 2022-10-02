package interfaces

//go:generate mockgen --destination=./mocks/table.go interfaces Table
type Table interface {
	SetTable(players ...Player)
	SpinTable()
	ShuffleTable()
	NumPlayers() int
	Seating() []Player
	SetSeating([]Player)
	CurrentTurn() int
	SetCurrentTurn(int)
	Seed() int64
	SetSeed(int64)
}
