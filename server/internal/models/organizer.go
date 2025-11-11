package models

import (
	"database/sql"
	"time"
)

type OrganizerModel struct {
	DB *sql.DB
}

type Organizer struct {
	Id               int       `json:"id"`
	UserId           int       `json:"user_id" binding:"required"`
	OrganizationName string    `json:"organization_name" binding:"required"`
	ContactEmail     string    `json:"contact_email"`
	CreatedAt        time.Time `json:"created_at"`
}
