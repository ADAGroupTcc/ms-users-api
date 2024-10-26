package exceptions

import "fmt"

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HandleExceptions(err error) ErrorResponse {
	internalErr, ok := err.(*Error)
	if !ok {
		return ErrorResponse{
			Code:    500,
			Message: "Internal server error",
		}
	}
	fmt.Println(internalErr.Error())

	parsedErr := internalErr.Err.(*Exception)
	return ErrorResponse{
		Code:    parsedErr.Code,
		Message: internalErr.Err.Error(),
	}
}
