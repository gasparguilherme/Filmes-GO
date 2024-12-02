package handlers

import (
	"errors"
)

func validateUser(name, email string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}

	if email == "" {
		return errors.New("email cannot be empty")
	}

	return nil
}
