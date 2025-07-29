package film

import (
	"context"
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/entities"
)

func (r Repository) SaveFilm(data entities.Film) (int, error) {
	query := `
		INSERT INTO films (title, director, year, gender, userID)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id;
	`

	var id int
	err := r.connectionInstance.QueryRow(context.TODO(), query,
		data.Title, data.Director, data.Year, data.Genre, data.UserID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("executando query: %w", err)
	}

	return id, nil
}
