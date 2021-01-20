-include .env
export

run:
	@go run ./cmd/api/main.go

test:
	@go test ./... -v

image:
	@docker build -t go-challenge-image .

docker:
	@docker-compose up --build app 