People API
This project is a REST API developed in Go that retrieves and enriches information about people (name, surname, patronymic, age, gender, nationality) using external APIs and stores the enriched data in a PostgreSQL database. It also provides various API endpoints for adding, updating, deleting, and querying people.

Features
Fetch people information by ID.

Add new people with enriched data (age, gender, nationality).

Delete people records.

Pagination and filtering for listing people.

Swagger documentation for the API.

PostgreSQL database integration.

Logging for debugging and information purposes.

Prerequisites
Before you start, make sure you have the following installed:

Go (version 1.18 or higher)

PostgreSQL (for the database)

Swagger for API documentation generation

(Optional) golangci-lint for linting

Installation
1. Clone the repository
bash
Копировать
Редактировать
git clone https://github.com/yourusername/people-api.git
cd people-api
2. Install dependencies
To install Go dependencies:

bash
Копировать
Редактировать
go mod tidy
3. Set up PostgreSQL
Make sure PostgreSQL is running and create a database.

bash
Копировать
Редактировать
# Connect to PostgreSQL
psql -U postgres
# Create the database
CREATE DATABASE people_db;
Ensure your .env file has the correct PostgreSQL connection details.

Example .env file:

ini
Копировать
Редактировать
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=people_db
4. Generate Swagger documentation
Generate the Swagger docs by running the following command:

bash
Копировать
Редактировать
swag init -g cmd/main.go
Running the Project
To start the server, run:

bash
Копировать
Редактировать
make run
This will build the application and start the API server on http://localhost:8080.

API Endpoints
1. Get all people
http
Копировать
Редактировать
GET /people
Description: Get a list of all people with optional pagination and filtering.

Query Params:

page: Page number (optional)

limit: Number of items per page (optional)

2. Get person by ID
http
Копировать
Редактировать
GET /people/{id}
Description: Get detailed information about a person by their ID.

3. Create a new person
http
Копировать
Редактировать
POST /people
Description: Add a new person to the database.

Request Body:

json
Копировать
Редактировать
{
  "name": "Dmitriy",
  "surname": "Ushakov",
  "patronymic": "Vasilevich"  // optional
}
4. Delete a person
http
Копировать
Редактировать
DELETE /people/{id}
Description: Delete a person by ID.

Testing
To run the tests for the project:

bash
Копировать
Редактировать
make test
Swagger UI
You can access the Swagger documentation for the API at:

http
Копировать
Редактировать
http://localhost:8080/swagger/index.html
Makefile Commands
Here are the available commands in the Makefile:

make build - Build the project

make run - Run the project

make test - Run tests

make lint - Lint the project (requires golangci-lint)

make fmt - Format the code

make clean - Clean the project (remove build artifacts)

make swagger - Generate Swagger documentation

make help - Show the available commands

Contributing
Feel free to fork the repository, create a new branch, and submit a pull request. Contributions are welcome!

License
This project is licensed under the MIT License - see the LICENSE file for details.

This README should give you a good overview of the project and help you run it successfully. Let me know if you need to add or modify any sections! 😊
