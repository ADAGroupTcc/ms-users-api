package exceptions

import (
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HandleExceptions(err error) ErrorResponse {
	customErr, ok := err.(*Error)
	if !ok {
		return ErrorResponse{
			Code:    500,
			Message: "Internal server error",
		}
	}

	switch customErr.Err {
	case ErrUserNotFound:
		return ErrorResponse{
			Code:    http.StatusNotFound,
			Message: customErr.Err.Error(),
		}
	case ErrInvalidPayload:
		return ErrorResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: customErr.Err.Error(),
		}
	case
		ErrInvalidFirstName,
		ErrInvalidLastName,
		ErrInvalidEmail,
		ErrInvalidCPF,
		ErrInvalidCategories,
		ErrUserAlreadyExists,
		ErrInvalidID:
		return ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: customErr.Err.Error(),
		}
	case ErrDatabaseFailure:
		return ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: customErr.Err.Error(),
		}
	default:
		return ErrorResponse{
			Code:    500,
			Message: "Internal server error",
		}
	}
}
