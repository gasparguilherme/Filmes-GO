package film

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func (h Handler) GetFilmsByUserID(w http.ResponseWriter, r *http.Request) {

	// acessa o query param "user_id" da URL
	userID := r.URL.Query().Get("user_id")

	// se não tiver o parâmetro ou ele for vazio, retorna um erro
	if userID == "" {
		slog.Error("user_id não fornecido na query", "error", "user_id obrigatorio")
		http.Error(w, "user_id é obrigatório", http.StatusBadRequest)
		return
	}

	// converte o userID para int
	userID_int, err := strconv.Atoi(userID)
	if err != nil {
		slog.Error("não foi possível converter user_id para inteiro", "user_id", userID_int, "error", err)
		http.Error(w, "user_id inválido", http.StatusBadRequest)
		return
	}

	err = validate.ValidateID(userID_int)
	if err != nil {
		slog.Error("id invalido", "user_id", userID_int, "error", err)
		http.Error(w, "por favor insira um id valido", http.StatusBadRequest)
		return
	}

	film, err := h.usecase.FindFilmsByUserID(userID_int)
	if err != nil {
		slog.Error("erro ao buscar filmes de usuario", "user_id", userID_int, "error", err)
		http.Error(w, "algo deu errado :(", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(film)
	if err != nil {
		slog.Error("erro ao converter para formato JSON", "user_id", userID_int, "error", err)
		return
	}
}
