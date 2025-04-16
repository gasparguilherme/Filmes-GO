package film

var idFilms int

func (Repository) GetNextFilmID() int {
	idFilms++
	return idFilms
}
