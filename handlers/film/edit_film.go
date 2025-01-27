package film

import (
	"encoding/json"
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/domain/usecases"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func editFilm(jsonInput []byte) {

	// minha duvida com essa estrura é em relação ao UserID, ele esta tendo alguma conexão com o campo ID na struct User
	var inputData struct {
		UserID int           `json:"user_id"`
		Film   entities.Film `json:"film"`
	}
	err := json.Unmarshal(jsonInput, &inputData)
	if err != nil {
		fmt.Println("error unmarshal json", err)
		return
	}

	err = validate.ValidateFilm(inputData.Film.Title, inputData.Film.Director, inputData.Film.Year,
		inputData.Film.Gender)
	if err != nil {
		fmt.Println(err)
		return
	}

	film := usecases.EditFilm(inputData.UserID, inputData.Film.ID, inputData.Film.Title, inputData.Film.Director, inputData.Film.Year, inputData.Film.Gender)

	jsonFilm, err := json.Marshal(film)
	if err != nil {
		fmt.Println("error converting json format ", err)
		return
	}

	fmt.Println("film successfully edited", string(jsonFilm))

}
