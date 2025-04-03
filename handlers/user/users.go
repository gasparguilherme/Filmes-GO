package user

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user *entities.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		slog.Error("não foi possivel interpretar o JSON", "error", err)
		return
	}

	err = validate.ValidateUser(user.Name, user.Email)
	if err != nil {
		slog.Error("error ao validar usuario", "error", err)
		return
	}

	user = user.NewUser(user.Name, user.Email)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
		return
	}

}
