package user

import (
	"encoding/json"
	"log/slog"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/domain/usecases"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func CreateUser(jsonInput []byte) {
	slog.Info("requisição de criação de usuario", "JSON", string(jsonInput))
	var user *entities.User
	err := json.Unmarshal(jsonInput, &user)
	if err != nil {
		slog.Error("não foi possivel interpretar o JSON", "error", err)
		return
	}

	err = validate.ValidateUser(user.Name, user.Email)
	if err != nil {
		slog.Error("error ao validar usuario", "error", err)
		return
	}

	user = usecases.NewUser(user.Name, user.Email)

	jsonUser, err := json.Marshal(user)
	if err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
		return
	}

	slog.Info("usuario criado com sucesso", "usuario", string(jsonUser))

}
