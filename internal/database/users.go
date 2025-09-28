package database

import (
	"context"
	"database/sql"
	"time"
)

type UsersModel struct {
	DB *sql.DB
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
}


func (m UsersModel) Insert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background() , 3*time.Second)
	defer cancel()

	query := `INSERT INTO users (email, password, name) VALUES ($1, $2, $3) RETURNING id`
	err := m.DB.QueryRowContext(ctx,query,user.Email, user.Password, user.Name).Scan(&user.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m UsersModel) Get(id int) (*User,error) {
	ctx, cancel := context.WithTimeout(context.Background() , 3*time.Second)
	defer cancel()

	var user User
	query := `SELECT * FROM users WHERE id = $1`

	err := m.DB.QueryRowContext(ctx, query, id).Scan(&user.Id,&user.Email,&user.Name,&user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil,nil
		}
		return nil,err
	}

	return &user, nil

}