package database

import (
	"database/sql"
	"log"
)

// global access through the project
var DB *sql.DB

// very simple DB to execute and create the database
const schemaString string = `
	CREATE TABLE IF NOT EXISTS submissions (
	id INTEGER NOT NULL PRIMARY KEY,
	was_valuable INTEGER NOT NULL,
	recommend INTEGER NOT NULL,
	sequel INTEGER NOT NULL,
	pacing INTEGER NOT NULL,
	comments TEXT
  );`

func InitDB() error {
	var path string = "./gohttp.DB"
	var err error

	// open the database
	DB, err = sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("database couldnt init: %v\n%v", err, DB)
		return err
	}

	// testing the connection
	if err = DB.Ping(); err != nil {
		log.Fatalf("database connection failed: %v", err)
		return err
	}

	_, err = DB.Exec(schemaString)
	if err != nil {
		log.Fatalf("database schema execution err: %v\n", err)
		return err
	}

	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
