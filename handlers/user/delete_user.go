package user

import (
	"log/slog"
	"net/http"
	"strconv"
)

func (h Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	rawID := r.PathValue("id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.usecase.DeleteUser(id)
	if err != nil {
		slog.Error("erro ao deletar usuário", "id", id, "error", err)
		http.Error(w, "erro ao deletar usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
