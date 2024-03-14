package tools

import (
	"fmt"
)

type FieldError struct {
	Field string
	Msg   string
}

// Method that returns the error message.
func (f *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", f.Field, f.Msg)
}

// Validate that the field is not empty.
func CheckField(fields map[string]any, requiredFields ...string) (err error) {
	for _, field := range requiredFields {
		if _, ok := fields[field]; !ok {
			return &FieldError{
				Field: field,
				Msg:   "is required",
			}
		}
	}
	return
}
