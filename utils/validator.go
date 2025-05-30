package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

var customMessages = map[string]map[string]string{
	"username": {
		"required": "username is required",
		"min":      "username must be at least 3 characters",
		"max":      "username must be less than 20 characters",
	},
	"email": {
		"required": "email is required",
		"email":    "email must be a valid email address",
	},
	"password": {
		"required": "Password is required",
		"min":      "Password must be at least 6 characters",
	},
}

func ValidateStruct(s interface{}) map[string]string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		field := e.Field()
		tag := e.Tag()

		// Look for a custom message
		if fieldMsgs, ok := customMessages[field]; ok {
			if msg, ok := fieldMsgs[tag]; ok {
				errors[field] = msg
				continue
			}
		}

		// Fallback generic message
		errors[field] = fmt.Sprintf("%s is not valid (%s)", field, tag)
	}

	return errors
}
