package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/samvel333/gorest/internal/models"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

// Create Person function
func (r *Repository) CreatePerson(person models.Person) error {
	query := `INSERT INTO people (name, surname, patronymic, age, gender, nationality)
	          VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.DB.Exec(query, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality)
	if err != nil {
		log.Println("Error", err)
	}
	return err
}
