-include .env
export

run:
	@go run ./cmd/api/main.go


image:
	@docker build -t go-challenge-image .

docker:
	@docker-compose up --build app 