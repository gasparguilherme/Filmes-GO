package film

import (
	"context"
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/entities"
)

func (r Repository) UpdateFilm(data entities.Film) error {
	query := `
		UPDATE films
		SET 
			title = $1, 
			director = $2,
			year = $3,
			genre = $4,
			userID = $5
		WHERE id = $6;
	`

	_, err := r.connectionInstance.Exec(
		context.TODO(),
		query,
		data.Title,
		data.Director,
		data.Year,
		data.Genre,
		data.UserID,
		data.ID,
	)

	if err != nil {
		return fmt.Errorf("erro ao atualizar filme: %w", err)
	}

	return nil
}
