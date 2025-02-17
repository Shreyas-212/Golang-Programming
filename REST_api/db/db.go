package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "host=localhost port=5432 user=postgres password=shre202124@ dbname=Rest_api sslmode=disable"

	var err error

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}


	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot reach the database:", err)
	}

	fmt.Println("Connected to PostgreSQL successfully!")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	query := "SELECT now()" 
	var currentTime string
	err = DB.QueryRow(query).Scan(&currentTime)
	if err != nil {
		panic("Error executing query")
	}

	fmt.Println("Current time from PostgreSQL:", currentTime)

	CreateTables()
}

func CreateTables() {
	var createUsersTable = `
	CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        email TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL
	)`
	_, err := DB.Exec(createUsersTable)
	if err!= nil {
        panic("Error creating users table")
    }
	fmt.Println("Table 'users' created successfully!")

	var createEventsTable = `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime TIMESTAMP NOT NULL,
		user_id INTEGER REFERENCES users(id)
	)`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		log.Fatal("Error creating events table:", err)
	}

	fmt.Println("Table 'events' created successfully!")

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY SERIAL,
		event_id INTEGER REFERENCES events(id),
		user_id INTEGER REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registrations table.")
	}
}