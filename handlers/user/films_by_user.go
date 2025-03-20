package user

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gasparguilherme/my-repository/domain/usecases"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

type FilmByUserRequest struct {
	UserID int `json:"user_id"`
}

func HandleGetFilmsByUserID(w http.ResponseWriter, r *http.Request) {

	var User FilmByUserRequest

	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		slog.Error("não foi possivel interpretar o JSON", "error", err)
		return
	}

	err = validate.ValidateID(User.UserID)
	if err != nil {
		slog.Error("erro ao buscar usuario", "error", err)
		return
	}

	film, err := usecases.FilmByUser(User.UserID)
	if err != nil {
		slog.Error("por favor insira um valor maior que zero", "error", err)
		return
	}

	err = json.NewEncoder(w).Encode(film)
	if err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
		return
	}
}
