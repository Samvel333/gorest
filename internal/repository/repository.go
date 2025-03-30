package repository

import (
	"database/sql"
	"fmt"
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

// Get all persos with pagination
func (r *Repository) GetPeople(name, surname string, age int, limit, offset int) ([]models.Person, error) {
	query := `SELECT id, name, surname, patronymic, age, gender, nationality FROM people WHERE 1=1`
	args := []interface{}{}
	argID := 1

	if name != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", argID)
		args = append(args, "%"+name+"%")
		argID++
	}

	if surname != "" {
		query += fmt.Sprintf(" AND surname ILIKE $%d", argID)
		args = append(args, "%"+surname+"%")
		argID++
	}
	
	if age > 0 {
		query += fmt.Sprintf(" AND age = $%d", argID)
		args = append(args, age)
		argID++
	}

	query += fmt.Sprintf(" ORDER BY id LIMIT $%d OFFSET $%d", argID, argID+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var people []models.Person
	for rows.Next() {
		var person models.Person
		if err := rows.Scan(&person.ID, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender, &person.Nationality); err != nil {
			return nil, err
		}
		people = append(people, person)
	}
	return people, nil
}
