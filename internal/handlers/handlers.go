package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/samvel333/gorest/internal/models"
	"github.com/samvel333/gorest/internal/repository"
)

type Handler struct {
	Repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) CreatePersonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var person models.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, "JSON parsing error", http.StatusBadRequest)
		return
	}

	// Сохраняем в БД
	if err := h.Repo.CreatePerson(person); err != nil {
		http.Error(w, "Error storing in db", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}
