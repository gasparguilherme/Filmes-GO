package film

func (u Usecase) EditFilm(filmID int, newTitle, newDirector string, newYear int, newGender string) error {
	err := u.repository.UpdateFilm(filmID, newTitle, newDirector, newYear, newGender)
	return err

}
