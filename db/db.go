package db

import (
	"database/sql"
	"log"

	"github.com/fatih/color"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// ANSI color codes
const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "go_api.db")

	if err != nil {
		log.Fatal("Could not connect to database.")
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database ping failed:", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE, 
	password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUserTable)

	if err != nil {
		log.Fatal("Could not create users table!")
		return
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		log.Fatal("Could not create events table!")
		return
	}
	log.Println(color.GreenString("Database initialized successfully!"))
}
