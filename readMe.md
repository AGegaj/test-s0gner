# Test Signer

## Description

This project is a test signer application.

## Setup

1. Configure the `app.env` file for database connection:
    - Open the `app.env` file.
    - Set the database connection URL to `"postgresql://postgres:admin@localhost:5432/test_signer?sslmode=disable"`.
    - Save the file.

2. Install the golang migrate package:
    ```bash
    go get -u github.com/golang-migrate/migrate/v4/cmd/migrate
    ```

3. Run the migration command to migrate tables:
    ```bash
    migrate -path db/migrations -database "postgresql://postgres:admin@localhost:5432/test_signer?sslmode=disable" -verbose up
    ```

4. Install the viper package for loading configs:
    ```bash
    go get github.com/spf13/viper
    ```

5. Create Golang Code with SQLC:
    - Install sqlc https://docs.sqlc.dev/en/latest/overview/install.html
    - run sqlc generate

## Testing

To test the REST APIs, use the provided Postman collection file `REST API.postman_collection.json`. Import the collection into Postman and start testing the APIs.
