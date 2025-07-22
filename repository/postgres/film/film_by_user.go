package film

import (
	"context"
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/entities"
)

func (r Repository) FindFilmsByUserID(userID int) ([]entities.Film, error) {
	query := `
		SELECT id, title, director, year, genre, user_id
		FROM filmes
		WHERE user_id = $1;
	`

	rows, err := r.connectionInstance.Query(context.TODO(), query, userID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar filmes: %w", err)
	}
	defer rows.Close()

	var filmes []entities.Film

	for rows.Next() { //entender melhor depois
		var f entities.Film
		err := rows.Scan(&f.ID, &f.Title, &f.Director, &f.Year, &f.Genre, &f.User_ID)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear filme: %w", err)
		}
		filmes = append(filmes, f)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro nos resultados: %w", err)
	}

	return filmes, nil
}
