package user

import (
	"context"
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/entities"
)

func (r Repository) SaveUser(data entities.User) (int, error) {
	query := `
	INSERT INTO users(name, email)
	VALUES(
		$1, 
		$2
	)
	RETURNING id;
	`

	var id int
	err := r.conn.QueryRow(context.TODO(), query, data.Name, data.Email).Scan(id)
	if err != nil {
		return 0, fmt.Errorf("executando query: %w", err)

	}
	return id, nil
}
