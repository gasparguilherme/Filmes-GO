package handlers

import (
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/usecases"
)

func createFilm(title string, director string, year int, gender string) {

	err := validateFilm(title, director, year, gender)
	if err != nil {
		fmt.Println(err)
	}

	film := usecases.CreateFilm(title, director, year, gender)
	fmt.Println("film successfully created!", film)

}
