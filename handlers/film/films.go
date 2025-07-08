package film

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func (h Handler) CreateFilm(w http.ResponseWriter, r *http.Request) {

	var requestFilm entities.Film
	err := json.NewDecoder(r.Body).Decode(&requestFilm)
	if err != nil {
		slog.Error("não foi possivel interpretar json", "error", err)
		return
	}

	err = validate.ValidateFilm(requestFilm.Title, requestFilm.Director, requestFilm.Year, requestFilm.Genre)
	if err != nil {
		slog.Error("erro ao validar filme", "error", err)
		return
	}

	film := h.usecase.CreateFilm(requestFilm.UserID, requestFilm.Title, requestFilm.Director, requestFilm.Year, requestFilm.Genre)

	err = json.NewEncoder(w).Encode(film)
	if err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
		return
	}

}
