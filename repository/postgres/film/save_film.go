package film

import (
	"context"
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/entities"
)

func (r Repository) SaveFilm(data entities.Film, created_by int) (int, error) {
	query := `
	INSERT INTO films (title, director, year, genre, userID)
	VALUES ($1, 
			$2, 
			$3, 
			$4, 
			$5)
	RETURNING id;
	`
	var id int
	err := r.connectionInstance.QueryRow(context.TODO(), query, data.Title, data.Director, data.Year, data.Genre,
		data.UserID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("erro ao salvar filme: %w", err)
	}
	return id, nil
}
