package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecords          = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// tries to signup with an email address that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

type Users struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	CreatedAt      time.Time
}
