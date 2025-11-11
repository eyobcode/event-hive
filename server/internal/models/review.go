package models

import (
	"database/sql"
	"time"
)

type ReviewModel struct {
	DB *sql.DB
}

type Review struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id" binding:"required"`
	EventId   int       `json:"event_id" binding:"required"`
	Rating    int       `json:"rating" binding:"required"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
