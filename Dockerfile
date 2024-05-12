FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

FROM alpine:latest
WORKDIR /app
COPY --from=builder /main /main
EXPOSE 9000
CMD ["/main"]
