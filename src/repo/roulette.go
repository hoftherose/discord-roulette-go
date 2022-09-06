package repo

import (
	"fmt"

	u "github.com/holy-tech/discord-roulette/src"
)

func CreateTable() {
	users_table := InitTable
	_, err := Connection.Exec(users_table)
	u.CheckErr("Error executing query: %v", err)
	fmt.Println("Table created successfully!")
}
