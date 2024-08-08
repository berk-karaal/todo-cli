package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "todo.sqlite") // TODO: get db path from config
	if err != nil {
		return nil, err
	}
	return db, nil
}

func SetupTodoDatabase(db *sql.DB) error {

	// status column choices are N: Not started, IP: In progress, D: Done
	// createdAt column is unix timestamp
	const createTable = `
CREATE TABLE IF NOT EXISTS todos(
id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
name TEXT NOT NULL,
status TEXT CHECK( status IN ('N','IP','D') )   NOT NULL DEFAULT 'N',
createdAt INTEGER NOT NULL
);
`

	_, err := db.Exec(createTable)
	if err != nil {
		return err
	}

	return nil
}
