build:
	go mod tidy
	go mod download
	go build -o main .

docker:
	docker compose up -d

kill-docker:
	docker compose down
run:
	./main -d