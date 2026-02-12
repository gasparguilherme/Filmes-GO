package user

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gasparguilherme/my-repository/authentication"
	"github.com/gasparguilherme/my-repository/domain/entities"
	"github.com/gasparguilherme/my-repository/handlers/validate"
)

func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	//requestID := r.Header.Get("Request-ID")
	//sessionID := r.Header.Get("Session-ID")

	//logger := slog.With(
	//"handler", "CreateUser",
	//"request_id", requestID,
	//"session_id", sessionID,
	//"method", r.Method,
	//	"path", r.URL.Path)

	//logger.Info("iniciando criacao de usuario")

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

	//token
	token, err := authentication.CreateToken(createdUser.ID)
	if err != nil {
		slog.Error("erro ao gerar token", "error", err)
		http.Error(w, "erro ao gerar token", http.StatusInternalServerError)
		return
	}
	response := AuthResponse{
		User:  *createdUser,
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
		http.Error(w, "ocorreu um erro inesperado", http.StatusInternalServerError)
		return
	}

}
