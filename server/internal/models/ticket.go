package models

import (
	"database/sql"
	"time"
)

type TicketModel struct {
	DB *sql.DB
}

type Ticket struct {
	Id         int       `json:"id"`
	EventId    int       `json:"event_id" binding:"required"`
	TicketType string    `json:"ticket_type"`
	Price      int       `json:"price"`
	Quantity   int       `json:"quantity"`
	CreatedAt  time.Time `json:"created_at"`
}
