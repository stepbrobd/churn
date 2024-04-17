package validator

import (
	"errors"
	"strconv"
)

func IntConvertible(s string) error {
	if _, err := strconv.Atoi(s); err != nil {
		return errors.New("must be a number")
	}
	return nil
}

func IntConvertibleNullable(s string) error {
	if s == "" {
		return nil
	}
	return IntConvertible(s)
}
