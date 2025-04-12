FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY . ./
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/app.go
RUN go build -o /app/build/main

FROM alpine:latest AS production
COPY . ./
COPY --from=builder /app/build/main ./main
COPY --from=builder /app/docs ./docs
EXPOSE 8080
CMD [ "./main" ]