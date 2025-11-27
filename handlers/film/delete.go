package film

import (
	"net/http"
	"strconv"
)

func (h Handler) DeleteFilm(w http.ResponseWriter, r *http.Request) {
	filmIDParam := r.PathValue("id")
	userIDParam := r.PathValue("userID")

	id, err := strconv.Atoi(filmIDParam)
	if err != nil {
		http.Error(w, "ID invalido: o valor deve ser um número inteiro", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
		return
	}

	err = h.usecase.DeleteFilm(id, userID)
	if err != nil {
		http.Error(w, "error ao deletar filme", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
