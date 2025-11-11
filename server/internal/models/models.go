package models

import "database/sql"

type Models struct {
	Users      UserModel
	Reviews    ReviewModel
	Payments   PaymentModel
	Attendees  AttendeeModel
	Tickets    TicketModel
	Organizers OrganizerModel
	Categories CategoryModel
	Events     EventModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:      UserModel{DB: db},
		Reviews:    ReviewModel{DB: db},
		Payments:   PaymentModel{DB: db},
		Attendees:  AttendeeModel{DB: db},
		Tickets:    TicketModel{DB: db},
		Organizers: OrganizerModel{DB: db},
		Categories: CategoryModel{DB: db},
		Events:     EventModel{DB: db},
	}
}
