build:
	go mod tidy
	go mod download
	docker compose build

docker:
	docker compose up -d

stop:
	docker compose stop

down:
	docker compose down --volumes