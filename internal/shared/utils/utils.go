package utils

import (
	"meetup-app-hexa-arch/internal/shared/errors"
	"regexp"
	"time"
)

// ConvertToTimezone converts a time to the specified timezone.
func ConvertToTimezone(t time.Time, timezone string) (time.Time, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, errors.WrapError(err, "invalid timezone")
	}
	return t.In(location), nil
}

// FormatTime formats a time in a readable format.
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05 MST")
}

// ValidateEmail checks if an email string is in a valid format.
func ValidateEmail(email string) bool {
	// A simple regex for email validation
	re := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return MatchRegex(re, email)
}

// MatchRegex validates a string against a regular expression.
func MatchRegex(pattern, value string) bool {
	matched, err := regexp.MatchString(pattern, value)
	return err == nil && matched
}
