package film

import (
	"context"
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/entities"
)

func (r Repository) FindFilmByID(filmID int) (*entities.Film, error) {
	query := `
		SELECT id, title, director, year, genre, userID
		FROM filmes
		WHERE id = $1;
	`

	row := r.connectionInstance.QueryRow(context.TODO(), query, filmID)

	var film entities.Film
	err := row.Scan(&film.ID, &film.Title, &film.Director, &film.Year, &film.Genre, &film.UserID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar filme: %w", err)
	}

	return &film, nil
}
