package database

import (
	"database/sql"
	"log"
)

var db *sql.DB

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
	var path string = "./gohttp.db"
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("database couldnt init: %v\n%v", err, db)
	}

	_, err = db.Exec(schemaString)
	if err != nil {
		log.Fatalf("database schema execution err: %v\n", err)
	}

	defer db.Close()

	return nil
}
