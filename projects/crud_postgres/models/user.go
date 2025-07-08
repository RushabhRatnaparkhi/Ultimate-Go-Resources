package models

import (
	"database/sql"
	"errors"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(db *sql.DB, user *User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	return db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
}

func GetUserByID(db *sql.DB, id int) (*User, error) {
	user := &User{}
	query := `SELECT id, name, email FROM users WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return user, err
}

func UpdateUser(db *sql.DB, user *User) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
	_, err := db.Exec(query, user.Name, user.Email, user.ID)
	return err
}

func DeleteUser(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}
