package models

import (
	"time"
	"fmt"
	"REST_api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, datetime, user_id)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = db.DB.QueryRow(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID).Scan(&e.ID)
	if err != nil {
		fmt.Printf("Error executing query: %v", err)
		return fmt.Errorf("failed to insert event: %w", err) 
	}

	fmt.Printf("Event inserted successfully with ID: %d", e.ID)

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = $1"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event)Update() error{
	query := `
    UPDATE events
    SET name=$1, description=$2, location=$3, datetime=$4
    WHERE id=$5`
    _, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.ID)
    if err!= nil {
        fmt.Printf("Error executing query: %v", err)
    }
    fmt.Printf("Event updated successfully with ID: %d", event.ID)
	return nil
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id=$1"
	_, err := db.DB.Exec(query, event.ID)
	if err != nil {
        fmt.Printf("Error executing query: %v", err)
    }
	fmt.Printf("Event deleted successfully with ID: %d", event.ID)
	return nil
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES ($1, $2)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = $1 AND user_id = $2"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}