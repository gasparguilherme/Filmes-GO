package film

import (
	"encoding/json"
	"log/slog"

	"github.com/gasparguilherme/my-repository/domain/usecases"
	"github.com/gasparguilherme/my-repository/handlers/validate"
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

	err = usecases.EditFilm(inputData.UserID, inputData.Title, inputData.Director, inputData.Year, inputData.Gender)
	if err != nil {
		slog.Error("erro ao editar filme")
		return
	}

	slog.Info("filme editado com sucesso")

}
