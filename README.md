# ISI Transaction API
This is simple API for sumulate transaction saving

## Prerequisite
- Go 1.24
- PostgresSQL
- [Swago](https://github.com/swaggo/swag?tab=readme-ov-file)
- [goose](https://github.com/pressly/goose)
  
## Installation
- Clone this repository using commang `git clone git@github.com:fiqrikm18/ISI-SavingAPI.git`.
- Copy `.env.example` to `.env`.
- Fill the .env value for example
    ```
    DB_HOST=db
    DB_NAME=saving_transaction
    DB_USERNAME=postgres
    DB_PASSWORD=postgres
    DB_ENABLE_SSL=disable

    GOOSE_DRIVER=postgres
    GOOSE_DBSTRING=postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:5432/${DB_NAME}
    GOOSE_MIGRATION_DIR=./db/migrations
    GOOSE_TABLE=${DB_NAME}.migrations
    ```
- Run `go mod tidy` to download dependecies.
- Run `go install github.com/pressly/goose/v3/cmd/goose@latest` to install goose for database migration.
- Run `go install github.com/swaggo/swag/cmd/swag@latest` to install swago for generating API documentation.
- Run `goose up` to migrate database.
- Run `swag init -g cmd/app.go && go run main.go --port=<port_number> --host=<application_host>` to run application, or you can run `make dev`.
- Or you can use dev container using visual studio code.
- You can access swagger doc page by visiting `<base_url>` in this sample is `http://localhost:8080` or `http://localhost:8080/docs`.