package repository

var idFilms int

func GetNextFilmID() int {
	idFilms++
	return idFilms
}
