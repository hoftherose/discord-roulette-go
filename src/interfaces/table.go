package interfaces

//go:generate mockgen --destination=./../../mocks/table.go github.com/holy-tech/discord-roulette/src/interfaces Table
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
