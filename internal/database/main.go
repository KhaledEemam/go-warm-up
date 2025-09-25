package database

import "database/sql"

type Models struct {
	Users     UsersModel
	Attendees  AttendeeModel
	Events    EventsModel
}

func CreateDataBaseModels(db *sql.DB) Models {
	return Models{
		Users:     UsersModel{DB: db},
		Attendees: AttendeeModel{DB: db},
		Events:    EventsModel{DB: db},
	}
}
