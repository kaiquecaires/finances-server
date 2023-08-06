package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidationErrorsToString(errs validator.ValidationErrors) string {
	var result string
	for _, err := range errs {
		result += fmt.Sprintf("%s is %s\n", err.Field(), err.ActualTag())
	}
	return result
}
