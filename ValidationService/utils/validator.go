package utils

import (
	"Application/ValidationService/internal/domain"
	"errors"
)

const (
	badNameError    string = "Bad name: "
	badSurnameError string = "Bad surname: "
	badAgeError     string = "Bad age: "
	badIdError      string = "Bad id: "
)

func CreateValidation(user domain.User) error {
	if user.Name == "" {
		return errors.New(badNameError + "emtpy string")
	}

	if user.Surname == "" {
		return errors.New(badSurnameError + "emtpy string")
	}

	if user.Age <= 0 {
		return errors.New(badAgeError + "age is negative")
	}

	return nil
}
