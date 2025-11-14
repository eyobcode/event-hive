package models

import (
	"database/sql"
	"errors"
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

func (m *OrganizerModel) EnsureOrganizer(userId int) (int, error) {
	var id int
	err := m.DB.QueryRow("SELECT id FROM organizers WHERE user_id=$1", userId).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		// Some other database error
		return 0, err
	}
	return id, nil
}
