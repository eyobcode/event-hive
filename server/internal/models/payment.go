package models

import (
	"database/sql"
	"time"
)

type PaymentModel struct {
	DB *sql.DB
}

type Payment struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id" binding:"required"`
	TicketId  int       `json:"ticket_id" binding:"required"`
	Amount    float64   `json:"amount" `
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
