package repo

import (
	"database/sql"

	u "github.com/holy-tech/discord-roulette/src"
)

var Connection *sql.DB

func init() {
	var err error
	Connection, err = sql.Open("sqlite3", "file:locked.sqlite?cache=shared")
	u.CheckErr("Could not return DB connection: %v", err)
}
