package lib

import (
	"fmt"
	"strings"
)

// prompt user to confirm
func Confirm() bool {
	var input string
	fmt.Print("Are you sure you want to continue? (y/n): ")
	fmt.Scanln(&input)

	return strings.ToLower(input) == "y"
}
