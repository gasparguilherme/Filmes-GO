package film

import (
	"context"
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/entities"
)

func (r Repository) SaverFilm(data entities.Film) (int, error) {
	query := `
		INSERT INTO films (title, director, year, gender, created_by)
		VALUES(
			$1, 
			$2,
			$3,
			$4,
			$5
	)
	RETURNING id;	
	`
	var id int
	err := r.conn.QueryRow(context.TODO(), query, data.Title, data.Director, data.Year, data.Gender, data.UserID).Scan(id)
	if err != nil {
		return 0, fmt.Errorf("executando query: %w", err)

	}

	return id, nil

}
