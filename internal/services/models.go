package services

import (
	"database/sql"
)

type Models struct {
	Users     UserModel
	Events    EventModel
	Attendees AttendeeModel
}

func NewModels(database *sql.DB) Models {
	return Models{
		Users:     UserModel{DB: database},
		Events:    EventModel{DB: database},
		Attendees: AttendeeModel{DB: database},
	}
}
