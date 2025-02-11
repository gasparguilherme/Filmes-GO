package film

import (
	"encoding/json"
	"log/slog"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/domain/usecases"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func HandleGetFilmByID(jsonInput []byte) {
	slog.Info("requisição de busca de filme por ID", "JSON", string(jsonInput))
	var f entities.Film
	err := json.Unmarshal(jsonInput, &f)
	if err != nil {
		slog.Error("não foi possivel interpretar json", "error", err)
		return
	}

	err = validate.ValidateID(f.ID)
	if err != nil {
		slog.Error("erro ao buscar ID", "error", err)
		return
	}

	f, err = usecases.FilmByID(f.ID)
	if err != nil {
		slog.Error("por favor insira um valor maior que zero", "error", err)
		return
	}

	jsonFilm, err := json.Marshal(f)
	if err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
		return
	}

	slog.Info("filme encontrado", "filme", string(jsonFilm))
}
