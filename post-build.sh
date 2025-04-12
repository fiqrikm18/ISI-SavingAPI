go install github.com/pressly/goose/v3/cmd/goose@latest
go install github.com/swaggo/swag/cmd/swag@latest

source .env
goose up