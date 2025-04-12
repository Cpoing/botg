package models

import (
	"database/sql"
	"errors"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Authenticate(name, password string) (int, error) {
	var id int
	var storedPassword string

	stmt := "SELECT id, hashed_password FROM users WHERE name = ?"

	err := m.DB.QueryRow(stmt, name).Scan(&id, &storedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	if password != storedPassword {
		return 0, ErrInvalidCredentials
	}

	return id, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	var exists bool

	stmt := "SELECT EXISTS(SELECT true FROM users WHERE id = ?)"

	err := m.DB.QueryRow(stmt, id).Scan(&exists)
	return exists, err
}
