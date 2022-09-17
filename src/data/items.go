package data

type Item struct {
	Name        string
	Description string
	Inventory   int
	Effect      func(GameSettings)
}
