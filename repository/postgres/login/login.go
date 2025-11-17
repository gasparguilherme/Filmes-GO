package login

import (
	"context"
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/entities"
)

func (r Repository) FindUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	query := `SELECT id, name, email, password FROM users WHERE email = $1`

	err := r.connectionInstance.QueryRow(context.TODO(), query, email).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password,
	)
	if err != nil {
		return nil, fmt.Errorf("usuário não encontrado: %w", err)
	}

	return &user, nil
}
