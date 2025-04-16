package film

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func (h Handler) HandleEditFilm(w http.ResponseWriter, r *http.Request) {
	var inputData struct {
		UserID   int    `json:"user_id"`
		Title    string `json:"title"`
		Director string `json:"director"`
		Year     int    `json:"year"`
		Gender   string `json:"gender"`
	}
	err := json.NewDecoder(r.Body).Decode(&inputData)
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

	err = h.usecase.EditFilm(inputData.UserID, inputData.Title, inputData.Director, inputData.Year, inputData.Gender)
	if err != nil {
		slog.Error("erro ao editar filme", "error", err)
		return
	}

	slog.Info("filme editado com sucesso")

}
