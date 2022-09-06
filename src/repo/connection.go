package repo

import (
	"database/sql"

	u "github.com/holy-tech/discord-roulette/src"
	_ "github.com/lib/pq"
)

var Connection *sql.DB

func init() {
	var err error
	Connection, err = sql.Open("postgres", "postgres://root:root@database/test_db?sslmode=disable")
	u.CheckErr("Could not return DB connection: %v", err)
}
