package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(connStr string) {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	createHotelsTable := `
	CREATE TABLE IF NOT EXISTS hotels (
		room_id SERIAL PRIMARY KEY,
		hotel_name VARCHAR(255) NOT NULL,
		rate_per_night NUMERIC NOT NULL,
		max_guests INT NOT NULL,
		is_available BOOLEAN NOT NULL
	);`

	createAvailableDatesTable := `
	CREATE TABLE IF NOT EXISTS available_dates (
		id SERIAL PRIMARY KEY,
		hotel_id INT REFERENCES hotels(room_id) ON DELETE CASCADE,
		available_date DATE NOT NULL
	);`

	_, err = DB.Exec(createHotelsTable)
	if err != nil {
		log.Fatal("Failed to create hotels table:", err)
	}

	_, err = DB.Exec(createAvailableDatesTable)
	if err != nil {
		log.Fatal("Failed to create available_dates table:", err)
	}

	log.Println("Database initialized successfully")
}
