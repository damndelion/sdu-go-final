compose-up: ### Run docker-compose
	docker-compose up
.PHONY: compose-up

build-app: ### Build docker image of application
	docker build -t app .
.PHONY: build-app

swag-init: ### swag init
	swag init -g internal/controller/http/router.go
.PHONY: swag-init

linter-golangci: ### check by golangci linter
	golangci-lint run
.PHONY: linter-golangci



