package user

import "fmt"

func (u Usecase) DeleteUser(id int) error {
	err := u.repository.DeleteUser(id)
	if err != nil {
		return fmt.Errorf("falha em deletar usuario %w", err)

	}
	return nil
}
