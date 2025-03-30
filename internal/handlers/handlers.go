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

// @Summary Create a new person
// @Description Creates a new person by storing the provided JSON payload in the database. The person's age, gender, and nationality are enriched via external APIs.
// @Tags persons
// @Accept json
// @Produce json
// @Param person body models.CreatePerson true "Person object to create"
// @Success 201 {object} models.CreatePerson "Created person with enriched data"
// @Failure 400 {string} string "JSON parsing error"
// @Failure 500 {string} string "Error storing in db"
// @Router /person [post]
func (h *Handler) CreatePersonHandler(w http.ResponseWriter, r *http.Request) {
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

// @Summary Get persons list
// @Description Returns persons list with filters and pagination
// @Param name query string false "NAme"
// @Param surname query string false "Surname"
// @Param age query int false "Age"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} models.Person
// @Router /people [get]
func (h *Handler) GetPeopleHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	surname := query.Get("surname")
	age, _ := strconv.Atoi(query.Get("age"))
	limit, _ := strconv.Atoi(query.Get("limit"))
	offset, _ := strconv.Atoi(query.Get("offset"))

	if limit == 0 {
		limit = 10 // default value
	}

	people, err := h.Repo.GetPeople(name, surname, age, limit, offset)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

// @Summary Get a person by ID
// @Description Retrieves a person from the database using the provided ID.
// @Tags persons
// @Accept json
// @Produce json
// @Param id query string true "Person ID"
// @Success 200 {object} models.Person
// @Failure 404 {string} string "Person not found"
// @Failure 500 {string} string "Internal server error"
// @Router /person [get]
func (h *Handler) GetPersonByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	person, err := h.Repo.GetPersonByID(id)
	if err != nil {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

// @Summary Delete a person by ID
// @Description Deletes a person from the database using the provided ID.
// @Tags persons
// @Accept json
// @Produce json
// @Param id query string true "Person ID"
// @Success 200 {object} map[string]string "Person deleted!"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Error deleting"
// @Router /person [delete]
func (h *Handler) DeletePersonHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if err := h.Repo.DeletePerson(id); err != nil {
		http.Error(w, "Error deleting", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Person deleted!"})
}

// @Summary Update a person
// @Description Updates an existing person in the database using the provided JSON payload.
// @Tags persons
// @Accept json
// @Produce json
// @Param person body models.Person true "Person object to update"
// @Success 200 {object} map[string]string "updated!"
// @Failure 400 {string} string "Error JSON parsing"
// @Failure 500 {string} string "error to update"
// @Router /person [put]
func (h *Handler) UpdatePersonHandler(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, "Error JSON parsing", http.StatusBadRequest)
		return
	}

	if err := h.Repo.UpdatePerson(person); err != nil {
		http.Error(w, "error to update", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "updated!"})
}
