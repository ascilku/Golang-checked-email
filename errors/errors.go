package errors

import "github.com/go-playground/validator/v10"

func ErrorMessage(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
