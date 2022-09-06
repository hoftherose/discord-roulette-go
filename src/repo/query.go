package repo

const InitTable = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"user_name" TEXT,
		"num_wins" TEXT,
		"num_losses" TEXT,
		"created_at" INT
		"deleted_at" INT
	);`
