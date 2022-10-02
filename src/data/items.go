package data

import "fmt"

type Item struct {
	Name        string
	Description string
	Inventory   int
	Effect      func(*GameStatus)
}

func (i *Item) Use(s *GameStatus) string {
	i.Effect(s)
	return i.Name + " was used, " + fmt.Sprint(i.Inventory) + " uses left"
}
