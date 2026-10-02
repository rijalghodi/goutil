package validate

import "fmt"

// Positive returns an error if n is not greater than zero.
// This package lives under internal/, so only code in this module can import it.
func Positive(n int, name string) error {
	if n <= 0 {
		return fmt.Errorf("%s must be greater than 0", name)
	}
	return nil
}
