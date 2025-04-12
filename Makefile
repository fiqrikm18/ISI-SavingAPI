dev:
	swag init -g cmd/app.go && go run main.go --port=8080 --host=localhost