.DEFAULT_GOAL := usage

usage:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build:
	docker compose build

build-prod:
	docker build -t gqlgen-demo-prod --target runner .

up:
	docker compose up

generate:
	go run github.com/99designs/gqlgen
