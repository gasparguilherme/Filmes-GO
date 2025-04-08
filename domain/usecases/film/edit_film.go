package film

import "github.com/gasparguilherme/my-repository/repository/film"

func (Usecase) EditFilm(filmID int, newTitle, newDirector string, newYear int, newGender string) error {
	filmRepository := film.Repository{}
	err := filmRepository.UpdateFilm(filmID, newTitle, newDirector, newYear, newGender)
	return err

}
