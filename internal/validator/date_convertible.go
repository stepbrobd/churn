package validator

import (
	"errors"
	"time"
)

func DateConvertible(s string) error {
	// convertible to time.Time in the format of "2006-01-02"
	if _, err := time.Parse("2006-01-02", s); err != nil {
		return errors.New("must be in the format of 'YYYY-MM-DD'")
	}
	return nil
}

func DateConvertibleNullable(s string) error {
	if s == "" {
		return nil
	}
	return DateConvertible(s)
}
