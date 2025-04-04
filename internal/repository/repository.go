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
func (r *Repository) CreatePerson(person models.Person) (*models.Person, error) {
	
	query := `INSERT INTO people (name, surname, patronymic, age, gender, nationality)
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, surname, patronymic, age, gender, nationality`

	var createdPerson models.Person
	err := r.DB.QueryRow(query, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality).
		Scan(&createdPerson.ID, &createdPerson.Name, &createdPerson.Surname, &createdPerson.Patronymic, &createdPerson.Age, &createdPerson.Gender, &createdPerson.Nationality)

	if err != nil {
		log.Println("Error", err)
		return nil, err
	}

	return &createdPerson, nil
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

func (r *Repository) GetPersonByID(id string) (models.Person, error) {
	var person models.Person
	query := `SELECT id, name, surname, patronymic, age, gender, nationality FROM people WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&person.ID, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender, &person.Nationality)

	if err != nil {
		return models.Person{}, err
	}
	return person, nil
}

// Delete person by id
func (r *Repository) DeletePerson(id string) error {
	query := `DELETE FROM people WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

// update method
func (r *Repository) UpdatePerson(person models.Person) error {
	query := `UPDATE people SET name=$1, surname=$2, patronymic=$3, age=$4, gender=$5, nationality=$6 WHERE id=$7`
	_, err := r.DB.Exec(query, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality, person.ID)
	return err
}
