package validate

import (
	"errors"
	"strings"
)

func ValidateFilm(title string, director string, year int, genre string) error {
	switch "" {
	case strings.TrimSpace(title):
		return errors.New("o titulo nao pode estar vazio")
	case strings.TrimSpace(director):
		return errors.New("o diretor nao pode estar vazio")
	case strings.TrimSpace(genre):
		return errors.New("o genero nao pode estar vazio")
	}
	if year <= 0 || year > 2024 {
		return errors.New("o ano deve ser maior que zero, e nao pode exceder 2024")

	}

	return nil

}
