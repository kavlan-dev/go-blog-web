package handlers

import (
	"encoding/json"
	"golang-blog-web/internal/models"
	"net/http"
)

type UserService interface {
	CreateUser(newUser *models.User) error
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "некорректный JSON", http.StatusBadRequest)
		return
	}

	newUser := &models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	if err := h.service.CreateUser(newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
