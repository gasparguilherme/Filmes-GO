package user

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest *entities.User
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		slog.Error("não foi possivel interpretar o JSON", "error", err)
		return
	}

	err = validate.ValidateUser(userRequest.Name, userRequest.Email)
	if err != nil {
		slog.Error("error ao validar usuario", "error", err)
		return
	}

	createdUser, err := h.usecase.NewUser(userRequest.Name, userRequest.Email)
	if err != nil {
		slog.Error("Erro ao criar usuario", "error", err)
		return
	}

	err = json.NewEncoder(w).Encode(createdUser)
	if err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
		return
	}

}
