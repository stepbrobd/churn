package validator

import "fmt"

func NotNull(s string) error {
	if s == "" {
		return fmt.Errorf("cannot be empty")
	}
	return nil
}
