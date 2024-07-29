package exceptions

import (
	"fmt"
)

const prefix = "users-api"

var (

	// Errors related to request validation
	ErrInvalidPayload    = fmt.Errorf("%s: invalid payload", prefix)
	ErrInvalidFirstName  = fmt.Errorf("%s: invalid first name", prefix)
	ErrInvalidLastName   = fmt.Errorf("%s: invalid last name", prefix)
	ErrInvalidEmail      = fmt.Errorf("%s: invalid email", prefix)
	ErrInvalidCPF        = fmt.Errorf("%s: invalid CPF", prefix)
	ErrInvalidCategories = fmt.Errorf("%s: invalid categories", prefix)
	ErrUserAlreadyExists = fmt.Errorf("%s: user already exists", prefix)
	ErrInvalidID         = fmt.Errorf("%s: invalid ID", prefix)

	// Database related errors
	ErrUserNotFound    = fmt.Errorf("%s: user not found", prefix)
	ErrDatabaseFailure = fmt.Errorf("%s: database failure", prefix)
)
