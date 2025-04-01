package film

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gasparguilherme/my-repository/domain/usecases"
	"github.com/gasparguilherme/my-repository/handlers/validate"
	"github.com/gorilla/mux"
)

func HandleGetFilmByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("erro ao converter id para inteiro", "error", err)
		http.Error(w, "ID invalido", http.StatusBadRequest)
	}

	err = validate.ValidateID(id)
	if err != nil {
		slog.Error("erro ao buscar ID", "error", err)
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return

	}

	film, err := usecases.FilmByID(id)
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
