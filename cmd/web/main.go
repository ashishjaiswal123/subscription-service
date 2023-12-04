package main

import (
	"database/sql"
	"log"
	"os"
)

const webPort = "80"

func main() {
	// connect to the database
	initDB()

	// create sessions

	// create channels

	// create waitgroup

	// setup the application config

	// set up mail

	// listen for web connection
}

func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("can't connect to database")
	}
}

func connectToDB() *sql.DB {
	counts := 0

	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
}
