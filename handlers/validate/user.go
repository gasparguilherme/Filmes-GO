package validate

import (
	"errors"
)

func ValidateUser(name, email string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}

	if email == "" {
		return errors.New("email cannot be empty")
	}

	return nil
}
