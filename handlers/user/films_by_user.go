package user

import (
	"encoding/json"
	"log/slog"

	"github.com/gasparguilherme/my-repository/handlers/validate"
	"github.com/gasparguilherme/my-repository/repository/user"
)

type FilmByUserRequest struct {
	UserID int `json:"user_id"`
}

func FilmByUser(jsonInput []byte) {
	slog.Info("requisição de busca de filme atraves de usuario", "JSON", string(jsonInput))

	var User FilmByUserRequest

	err := json.Unmarshal(jsonInput, &User)
	if err != nil {
		slog.Error("não foi possivel interpretar o JSON", "error", err)
		return
	}

	err = validate.ValidateID(User.UserID)
	if err != nil {
		slog.Error("erro ao buscar usuario", "error", err)
		return
	}

	u, err := user.FilmByUser(User.UserID)
	if err != nil {
		slog.Error("por favor insira um valor maior que zero", "error", err)
		return
	}

	jsonUser, err := json.Marshal(u)
	if err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
		return
	}
	slog.Info("filme encontrado", "filme", string(jsonUser))
}
