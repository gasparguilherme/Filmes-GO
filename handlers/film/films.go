package film

import (
	"encoding/json"
	"log/slog"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/domain/usecases"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func CreateFilm(jsonInput []byte) {
	slog.Info("requisição de criação de filme", "JSON", string(jsonInput))
	var requestFilm entities.Film
	err := json.Unmarshal(jsonInput, &requestFilm)
	if err != nil {
		slog.Error("não foi possivel interpretar json", "error", err)
		return
	}

	err = validate.ValidateFilm(requestFilm.Title, requestFilm.Director, requestFilm.Year, requestFilm.Gender)
	if err != nil {
		slog.Error("erro ao validar filme", "error", err)
		return
	}

	film := usecases.CreateFilm(requestFilm.UserID, requestFilm.Title, requestFilm.Director, requestFilm.Year, requestFilm.Gender)

	jsonFilm, err := json.Marshal(film)
	if err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
		return
	}

	slog.Info("filme criado com sucesso!", "filme", string(jsonFilm))

}
