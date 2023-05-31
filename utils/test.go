package utils

import "fmt"

// FormatTestError format for you your test error messages.
//
// excepted is the value excepted.
// got is the value got.
//
// Return the formatted string
func FormatTestError(msg string, excepted string, got string) string {
	return fmt.Sprintf("%s\nexcepted: %s\ngot %s\n", msg, excepted, got)
}
