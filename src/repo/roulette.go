package repo

import (
	"fmt"

	u "github.com/holy-tech/discord-roulette/src"
	_ "github.com/mattn/go-sqlite3"
)

func CreateTable() {
	users_table := InitTable
	query, err := Connection.Prepare(users_table)
	u.CheckErr("Error in preparing query: %v", err)
	_, err = query.Exec()
	u.CheckErr("Error executing query: %v", err)
	fmt.Println("Table created successfully!")
}
