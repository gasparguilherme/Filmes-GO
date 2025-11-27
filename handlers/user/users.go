package user

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest entities.User
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		slog.Error("não foi possivel interpretar o JSON", "error", err)
		http.Error(w, "ocorreu um erro inesperado", http.StatusInternalServerError)
		return
	}

	err = validate.ValidateUser(userRequest.Name, userRequest.Email)
	if err != nil {
		slog.Error("error ao validar usuario", "error", err)
		http.Error(w, "informações inválidas", http.StatusBadRequest)

		return
	}

	createdUser, err := h.usecase.NewUser(userRequest.Name, userRequest.Email, userRequest.Password)
	if err != nil {
		slog.Error("erro ao criar usuario", "error", err)
		http.Error(w, "erro ao criar usuario", http.StatusBadRequest)

		return
	}

	err = json.NewEncoder(w).Encode(createdUser)
	if err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
		http.Error(w, "ocorreu um erro inesperado", http.StatusInternalServerError)
		return
	}

}
