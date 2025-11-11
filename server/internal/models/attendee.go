package models

import (
	"database/sql"
	"time"
)

type AttendeeModel struct {
	DB *sql.DB
}

type Attendee struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id" binding:"required"`
	EventId   int       `json:"event_id" binding:"required"`
	CheckedID int       `json:"checked_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}
