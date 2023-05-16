# Define variables
GO_SCRIPT := "$(PWD)/cmd/generate_structs"
# DB_NAME=temp_$$(openssl rand -hex 8);
DB_NAME=temp

.PHONY: all create-db run-script delete-db

all: create-db apply-sql run-script delete-db
# all: current

create-db:
	# Generate a random database name
	echo "Random database name: $(DB_NAME)"; \
	psql -c "DROP DATABASE IF EXISTS $(DB_NAME)"
	psql -c "CREATE DATABASE $(DB_NAME)"

apply-sql:
	psql -d "$(DB_NAME)" -f queries.sql

run-script:
	# Get the DSN and pass it to the Golang script
	DSN="host=localhost user=postgres password=password dbname=$(DB_NAME) port=5432 sslmode=disable"; \
	cd $(GO_SCRIPT); \
	DSN="$$DSN" go run main.go;

deletedb:
	# Delete the database
	psql -c "DROP DATABASE IF EXISTS $(DB_NAME)"

current:
	echo $(PWD)