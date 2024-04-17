package validator

import (
	"errors"
	"strconv"
)

func FloatConvertible(s string) error {
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return errors.New("must be a number")
	}
	return nil
}

func FloatConvertibleNullable(s string) error {
	if s == "" {
		return nil
	}
	return FloatConvertible(s)
}
