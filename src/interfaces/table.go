package interfaces

//go:generate mockgen --destination=./mocks/table.go interfaces Table
type Table interface {
	SetTable(players ...User)
	SpinTable()
	ShuffleTable()
	NumPlayers() int
	Seating() []User
	SetSeating([]User)
	CurrentTurn() int
	SetCurrentTurn(int)
	Seed() int64
	SetSeed(int64)
}
