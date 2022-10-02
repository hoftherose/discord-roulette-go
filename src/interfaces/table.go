package interfaces

//go:generate mockgen --destination=./mocks/table.go interfaces Table
type Table interface {
	InitTable(players ...User)
	SpinTable()
	ShuffleTable()
	NumPlayers() int
	GetSeating() []User
	SetSeating([]User)
	GetCurrentTurn() int
	SetCurrentTurn(int)
	GetSeed() int64
	SetSeed(int64)
}
