package exceptions

import "fmt"

type Error struct {
	Err        error
	TraceError error
}

func New(err error, traceError error) *Error {
	return &Error{
		Err:        err,
		TraceError: traceError,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Err, e.TraceError)
}
