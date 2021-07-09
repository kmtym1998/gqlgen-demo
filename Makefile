.DEFAULT_GOAL := usage

usage:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build:
	docker build -t gqlgen-demo .

run:
	go run server.go

generate:
	go generate ./...

up:
	docker run -p 8081:8081 --name gqlgen-demo -it gqlgen-demo

