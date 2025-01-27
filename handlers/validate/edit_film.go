package validate

import (
	"errors"
	"strings"
)

func ValidateFilm(title string, director string, year int, gender string) error {
	switch "" {
	case strings.TrimSpace(title):
		return errors.New("the title cannot be empty")
	case strings.TrimSpace(director):
		return errors.New("the director cannot be empty")
	case strings.TrimSpace(gender):
		return errors.New("the gender cannot be empty")
	}
	if year <= 0 || year > 2024 {
		return errors.New("the year must be greater than zero and not exceed 2024")

	}

	return nil

}
