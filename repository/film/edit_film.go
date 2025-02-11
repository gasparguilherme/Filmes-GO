package film

import "errors"

func UpdateFilm(filmID int, newTitle string, newDirector string, newYear int, newGender string) error {
	films := GetFilms()
	for i, film := range films {
		if film.ID == filmID {
			films[i].Title = newTitle
			films[i].Director = newDirector
			films[i].Year = newYear
			films[i].Gender = newGender
			return nil

		}
	}
	return errors.New("filme não encontrado para o usuário")

}
