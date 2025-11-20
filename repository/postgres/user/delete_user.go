package user

import (
	"context"
	"fmt"
)

func (r Repository) DeleteUser(id int) error {
	query := `DELETE FROM users 
	WHERE id = $1`

	_, err := r.connectionInstance.Exec(context.TODO(), query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar usuario: %w", err)
	}

	return nil
}
