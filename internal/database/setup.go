package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

func DBLocation() string {
	dbLocation := viper.GetString("database.location")
	return dbLocation
}

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DBLocation())
	if err != nil {
		return nil, err
	}
	return db, nil
}

func DropTables(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS todos;")
	if err != nil {
		return err
	}
	return nil
}

// CreateTables create necessary database tables if not exists
func CreateTables(db *sql.DB) error {

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
