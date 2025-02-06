package film

var idFilms int = 1

func GetNextFilmID() int {
	idFilms++
	return idFilms
}
