package film

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func (h Handler) CreateFilm(w http.ResponseWriter, r *http.Request) {

	requestID := r.Header.Get("Request-ID")
	sessionID := r.Header.Get("Session-ID")

	logger := slog.With(
		"handler", "CreateUser",
		"request_id", requestID,
		"session_id", sessionID,
		"method", r.Method,
		"path", r.URL.Path)

	logger.Info("iniciando criacao de filme")

	var requestFilm entities.Film
	err := json.NewDecoder(r.Body).Decode(&requestFilm)
	if err != nil {
		slog.Error("não foi possivel interpretar json", "error", err)
		http.Error(w, "ocorreu um erro inesperado", http.StatusInternalServerError)
		return
	}

	err = validate.ValidateFilm(requestFilm.Title, requestFilm.Director, requestFilm.Year, requestFilm.Genre)
	if err != nil {
		slog.Error("erro ao validar filme", "error", err)
		http.Error(w, "informações inválidas", http.StatusBadRequest)
		return

	}

	film, err := h.usecase.CreateFilm(requestFilm.UserID, requestFilm.Title, requestFilm.Director, requestFilm.Year, requestFilm.Genre)
	if err != nil {
		slog.Error("erro ao criar filme", "error", err)
		http.Error(w, "ocorreu um erro inesperado", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(film)
	if err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
		http.Error(w, "ocorreu um erro inesperado", http.StatusInternalServerError)
		return
	}

}
