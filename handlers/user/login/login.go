package login

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gasparguilherme/my-repository/domain/usecases/user/login"
)

func (h Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginData login.Login

	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		slog.Error("não foi possivel interpretar o JSON", "error", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	userLogged, err := h.usecase.Login(loginData)
	if err != nil {
		slog.Error("dados de login inválido", "error", err)
		http.Error(w, "dados de login inválido", http.StatusUnauthorized)
		return
	}

	userLogged.Password = ""

	err = json.NewEncoder(w).Encode(userLogged)
	if err != nil {
		slog.Error("erro ao converter para JSON", "error", err)
		http.Error(w, "erro ao gerar resposta", http.StatusInternalServerError)
		return
	}
}
