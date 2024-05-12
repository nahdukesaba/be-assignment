FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go mod tidy
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g server/server.go
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM library/postgres
COPY db/init.sql /docker-entrypoint-initdb.d/

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 9000
CMD ["./main"]
