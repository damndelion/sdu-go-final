compose-up: ### Run docker-compose
	docker-compose up --build -d postgres rabbitmq && docker-compose logs -f
.PHONY: compose-up

compose-down: ### Down docker-compose
	docker-compose down --remove-orphans
.PHONY: compose-down

swag-v1: ### swag init
	swag init -g internal/controller/http/router.go
.PHONY: swag-v1

linter-golangci: ### check by golangci linter
	golangci-lint run
.PHONY: linter-golangci



