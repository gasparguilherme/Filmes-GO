package film

import (
	"context"
	"fmt"
)

func (r Repository) DeleteFilm(id, userID int) error {
	query := `
		DELETE FROM films
		WHERE id = $1 AND userID = $2
	`

	_, err := r.connectionInstance.Exec(context.TODO(), query, id, userID)
	if err != nil {
		return fmt.Errorf("executando query de deleção: %w", err)
	}

	return nil
}
