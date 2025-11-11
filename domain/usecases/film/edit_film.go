package film

import "github.com/gasparguilherme/my-repository/domain/entities"

func (u Usecase) EditFilm(filmID int, newTitle, newDirector string, newYear int, newGenre string, userID int) error {
	film := entities.Film{
		ID:       filmID,
		Title:    newTitle,
		Director: newDirector,
		Year:     newYear,
		Genre:    newGenre,
		UserID:   userID,
	}

	err := u.repository.UpdateFilm(film)
	return err
}
