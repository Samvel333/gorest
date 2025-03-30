package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/samvel333/gorest/internal/models"
	"github.com/samvel333/gorest/internal/repository"
	"github.com/samvel333/gorest/pkg/httpclient"
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

	age, err := httpclient.GetAge(person.Name)
	if err != nil {
		log.Println("Error getting age:", err)
	}
	gender, err := httpclient.GetGender(person.Name)
	if err != nil {
		log.Println("Error getting gender:", err)
	}
	nationality, err := httpclient.GetNationality(person.Name)
	if err != nil {
		log.Println("Error getting nationality:", err)
	}

	person.Age = age
	person.Gender = gender
	person.Nationality = nationality

	// Save in DB
	if err := h.Repo.CreatePerson(person); err != nil {
		http.Error(w, "Error storing in db", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

func (h *Handler) GetPeopleHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	surname := query.Get("surname")
	age, _ := strconv.Atoi(query.Get("age"))
	limit, _ := strconv.Atoi(query.Get("limit"))
	offset, _ := strconv.Atoi(query.Get("offset"))

	if limit == 0 {
		limit = 10 // значение по умолчанию
	}

	people, err := h.Repo.GetPeople(name, surname, age, limit, offset)
	if err != nil {
		http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}