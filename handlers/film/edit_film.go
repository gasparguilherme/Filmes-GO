package film

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/gasparguilherme/my-repository/handlers/validate"
	"github.com/gasparguilherme/my-repository/repository/film"
)

func EditFilm(jsonInput []byte) {
	slog.Info("requisição de edição de filme", "JSON", string(jsonInput))
	var inputData struct {
		UserID   int `json:"user_id"`
		Title    string
		Director string
		Year     int
		Gender   string
	}
	err := json.Unmarshal(jsonInput, &inputData)
	if err != nil {
		slog.Error("erro ao interpretar formato JSON", "error", err)
		return
	}

	err = validate.ValidateFilm(inputData.Title, inputData.Director, inputData.Year,
		inputData.Gender)
	if err != nil {
		slog.Error("erro ao validar filme", "error", err)
		return
	}

	film := film.EditFilm(inputData.UserID, inputData.Title, inputData.Director, inputData.Year, inputData.Gender)

	jsonFilm, err := json.Marshal(film)
	if err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
		return
	}

	fmt.Println("filme editado com sucesso", string(jsonFilm))

}
