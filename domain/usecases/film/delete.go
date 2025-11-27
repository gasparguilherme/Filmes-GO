package film

import "fmt"

func (u Usecase) DeleteFilm(id, userID int) error {
	err := u.repository.DeleteFilm(id, userID)
	if err != nil {
		return fmt.Errorf("falha em deletar usuario %w", err)

	}
	return nil
}
