package errors

import (
	"runtime/debug"

	"go.uber.org/zap"
)

type (
	Causer interface {
		Cause() error
	}

	Error struct {
		Code  int
		Msg   string
		cause error
	}
)

func n(msg string, code int, cause error) Error {
	return Error{
		Msg:   msg,
		Code:  code,
		cause: cause,
	}
}

// New returns a new error with the given message and code.
func New(msg string, code int) Error {
	return n(msg, code, nil)
}

// From returns a new error with the given message and code.
// The cause is set to the given error.
func From(err error, msg string, code int) error {
	return n(msg, code, err)
}

// CodeFrom returns the error code of the given error or fallbackCode.
func CodeFrom(err error, fallbackCode int) int {
	if e, ok := err.(Error); ok {
		return e.Code
	}
	return fallbackCode
}

// CodeIs returns true if the error code is the same as the given code.
func CodeIs(err error, code int) bool {
	if e, ok := err.(Error); ok {
		return e.Code == code
	}
	return false
}

func (e Error) Error() string {
	return e.Msg
}

// Cause returns the innermost cause of the error.
func (e Error) Cause() error {
	if e.cause != nil {
		if causer, ok := e.cause.(Causer); ok {
			innerCause := causer.Cause()
			if innerCause == nil {
				return e.cause
			}
			return innerCause
		}
		return e.cause
	}

	return nil
}

// ErrorLogFields returns the error fields for logging.
func ErrorLogFields(err error) []zap.Field {
	var fields []zap.Field

	switch e := err.(type) {
	case Error:
		fields = append(fields, zap.String("error", e.Msg))
		fields = append(fields, zap.Int("errorCode", e.Code))

		if e.cause != nil {
			fields = append(fields, zap.String("errorCause", e.Cause().Error()))
		}

	default:
		fields = append(fields, zap.String("error", e.Error()))
	}

	return append(fields, zap.String("stacktrace", string(debug.Stack())))
}
