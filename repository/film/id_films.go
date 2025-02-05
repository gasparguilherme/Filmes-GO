package film

var idFilms int

func GetNextFilmID() int {
	idFilms++
	return idFilms
}
