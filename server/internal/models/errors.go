package models

import (
	"errors"
)

var (
	ErrNoRecord = errors.New("Models: no matching record found")

	ErrInvalidCredentials = errors.New("Models: invalid credentials")
)
