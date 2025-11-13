package models

import (
	"context"
	"database/sql"
	"time"
)

type EventModel struct {
	DB *sql.DB
}

type Event struct {
	Id          int       `json:"id"`
	OrganizerId int       `json:"organizer_id" binding:"required"`
	CategoryId  int       `json:"category_id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedAt   time.Time `json:"created_at"`
}

func (m *EventModel) Insert(event *Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		INSERT INTO events (organizer_id, category_id, title, description, location, start_time, end_time)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at
	`

	// QueryRowContext uses ctx to enforce timeout or cancellation
	return m.DB.QueryRowContext(
		ctx,
		query,
		event.OrganizerId,
		event.CategoryId,
		event.Title,
		event.Description,
		event.Location,
		event.StartTime,
		event.EndTime,
	).Scan(&event.Id, &event.CreatedAt)
}
