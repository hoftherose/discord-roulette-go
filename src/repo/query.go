package repo

const InitTable = `
	DROP TABLE IF EXISTS users;
	CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		user_name VARCHAR (100),
		num_wins INT,
		num_losses INT,
		created_at TIMESTAMP WITH TIME ZONE,
		deleted_at TIMESTAMP WITH TIME ZONE
	);`
