package interfaces

//go:generate mockgen --destination=./mocks/table_state.go interfaces TableState
type TableState interface {
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
