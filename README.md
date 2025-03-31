

# People API
  
This project is a REST API developed in Go that retrieves and enriches information about people (name, surname, patronymic, age, gender, nationality)
 using external APIs and stores the enriched data in a PostgreSQL database. It also provides various API endpoints for adding, updating, deleting, and querying people.

## Features  

- Fetch people information by ID.
- Add new people with enriched data (age, gender, nationality).
- Delete people records. 
- Pagination and filtering for listing people.
- Swagger documentation for the API.
- PostgreSQL database integration.
- Logging for debugging and information purposes.


## Prerequisites
Before you start, make sure you have the following installed:

- Go (version 1.18 or higher)
- PostgreSQL (for the database) (You can use Docker image)
- Migrate CLI (Go migrate)
- Makefile CLI (for windows)


## Migrate CLI Setup

1. Install migrate package for postgres

~~~bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
~~~

2. Set go path manually in your project

~~~bash
export PATH=$PATH:$(go env GOPATH)/bin
~~~

3. Migrate CLI is installed. You can check!

~~~bash
migrate -v
~~~

## Run Project  

1. Clone the project  

~~~bash  
git clone https://github.com/Samvel333/gorest
~~~

2. Go to the project directory  

~~~bash  
cd gorest
~~~

3. Set up PostgreSQL
~~~bash  
# Connect to PostgreSQL
psql -U postgres
# Create the database
CREATE DATABASE {name}

# note: same db name you should use in your .env file;
~~~

Don't forget to create .env file. You can use .env-example template
~~~bash  
# Linux commands 
touch .env
cat .env-example
# note: Before next step you should have env file
~~~

4. Install dependencies  

~~~bash  
make install
~~~

5. Run migrations
~~~bash
make up
~~~

5. Generate swagger docs

~~~bash  
make swagger
~~~

6. Start the server  

~~~bash  
make run
~~~

7. You can use all in one command after 4th step
~~~bash
make start
~~~


## Environment Variables  

To run this project, you will need to add the following environment variables to your .env file  
`HOST` - Your host (localhost if you run locally)

`PORT` - Port of Base URL

`DB_HOST` - Host of pg database server

`DB_PORT` - Port of pg database server

`DB_USER` - Your db server username

`DB_PASSWORD` - Password of your db server

`DB_NAME` - name of database


## License  

[MIT](https://choosealicense.com/licenses/mit/)
