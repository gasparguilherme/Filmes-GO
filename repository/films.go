package repository

var id_films int

func GetNextID_films() int {
	id_films++
	return id_films
}
