package film

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/domain/usecases"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func HandleGetFilmByID(w http.ResponseWriter, r *http.Request) {
	var f entities.Film
	err := json.NewDecoder(r.Body).Decode(&f)
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
		slog.Error("falha ao buscar filme", "error", err)
		return
	}

	err = json.NewEncoder(w).Encode(f)
	if err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
		return
	}

}
