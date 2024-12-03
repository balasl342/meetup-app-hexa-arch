package errors

import (
	"errors"
	"fmt"
)

// Common errors
var (
	ErrNotFound          = errors.New("resource not found")
	ErrUnauthorized      = errors.New("unauthorized access")
	ErrInvalidInput      = errors.New("invalid input provided")
	ErrInternalServer    = errors.New("internal server error")
	ErrDatabaseOperation = errors.New("database operation failed")
)

// WrapError adds additional context to an error.
func WrapError(err error, message string) error {
	return fmt.Errorf("%s: %w", message, err)
}

// IsErrorType checks if an error is of a specific type.
func IsErrorType(err, target error) bool {
	return errors.Is(err, target)
}
