package database

import (
	"context"
	"database/sql"
	"time"
)

type AttendeeModel struct {
	DB *sql.DB
}

type Attendee struct {
	Id      int `json:"id"`
	UserId  int `json:"userId"`
	EventId int `json:"eventId"`
}

func (m AttendeeModel) GetByEventAndAttendee(eventId int, attendeeId int) (*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var existingAttendee Attendee
	query := `SELECT * FROM attendees WHERE event_id = $1 AND user_id = $2`

	err := m.DB.QueryRowContext(ctx, query, eventId, attendeeId).Scan(&existingAttendee.Id, &existingAttendee.UserId, &existingAttendee.EventId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &existingAttendee, nil
}

func (m AttendeeModel) Insert(attendee *Attendee) (*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO attendees (event_id, user_id) VALUES ($1, $2) RETURNING id`

	err := m.DB.QueryRowContext(ctx, query, attendee.EventId, attendee.UserId).Scan(&attendee.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return attendee, nil
}

func (m AttendeeModel) GetAttendeesByEvent(eventId int) ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
     SELECT u.id, u.name, u.email
     FROM users u
     JOIN attendees a ON u.id = a.user_id
     WHERE a.event_id = $1
 `
	rows, err := m.DB.QueryContext(ctx, query, eventId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	users := []User{}

	for rows.Next() {
		var attendee User

		err := rows.Scan(&attendee.Id, &attendee.Name, &attendee.Email)

		if err != nil {
			return nil, err
		}

		users = append(users, attendee)
	}

	return users, nil
}

func (m AttendeeModel) Delete(userId int, eventId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE FROM attendees WHERE user_id = $1 AND event_id = $2`
	_, err := m.DB.ExecContext(ctx, query, userId, eventId)
	if err != nil {
		return err
	}

	return nil
}
