package film

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func (h Handler) GetFilmByID(w http.ResponseWriter, r *http.Request) {

	rawID := r.PathValue("id")

	id, err := strconv.Atoi(rawID)
	if err != nil {
		slog.Error("erro ao converter id para inteiro", "id", rawID, "error", err)
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	err = validate.ValidateID(id)
	if err != nil {
		slog.Error("erro ao buscar ID", "error", err)
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return

	}

	film, err := h.usecase.FindFilmByID(id)
	if err != nil {
		slog.Error("erro ao buscar filme", "error", err)
		http.Error(w, "falha ao buscar filme", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(film)
	if err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
		http.Error(w, "erro interno", http.StatusInternalServerError)
		return
	}

}
