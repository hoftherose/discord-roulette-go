package repo

const InitTable = `
	CREATE TABLE users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"FirstName" TEXT,
		"LastName" TEXT,
		"Dept" TEXT,
		"Salary" INT
	);`
