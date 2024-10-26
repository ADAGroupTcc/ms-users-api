package exceptions

import (
	"errors"
)

const prefix = "users-api"

type Exception struct {
	Err    error
	Code   int
	Prefix string
}

func (e *Exception) Error() string {
	return e.Prefix + ": " + e.Err.Error()
}

var (

	// Errors related to request validation
	ErrInvalidPayload    = &Exception{Err: errors.New("invalid payload"), Code: 400, Prefix: prefix}
	ErrInvalidFirstName  = &Exception{Err: errors.New("invalid first name"), Code: 400, Prefix: prefix}
	ErrInvalidLastName   = &Exception{Err: errors.New("invalid last name"), Code: 400, Prefix: prefix}
	ErrInvalidEmail      = &Exception{Err: errors.New("invalid email"), Code: 400, Prefix: prefix}
	ErrInvalidCPF        = &Exception{Err: errors.New("invalid CPF"), Code: 400, Prefix: prefix}
	ErrInvalidCategories = &Exception{Err: errors.New("invalid categories"), Code: 400, Prefix: prefix}
	ErrUserAlreadyExists = &Exception{Err: errors.New("user already exists"), Code: 400, Prefix: prefix}
	ErrInvalidID         = &Exception{Err: errors.New("invalid ID"), Code: 400, Prefix: prefix}

	// Database related errors
	ErrUserNotFound    = &Exception{Err: errors.New("user not found"), Code: 404, Prefix: prefix}
	ErrDatabaseFailure = &Exception{Err: errors.New("database failure"), Code: 500, Prefix: prefix}
)
