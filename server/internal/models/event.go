package models

import (
	"database/sql"
	"time"
)

type EventModel struct {
	DB *sql.DB
}

type Event struct {
	Id          int        `json:"id"`
	OrganizerId int        `json:"organizer_id" binding:"required"`
	CategoryId  int        `json:"category_id" binding:"required"`
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description"`
	Location    string     `json:"location"`
	StartTime   time.Time  `json:"start_time"`
	EndTime     time.Timer `json:"end_time"`
}
