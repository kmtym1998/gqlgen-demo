.DEFAULT_GOAL := usage

usage:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build:
	docker build -t gqlgen-demo .

up:
	go run server.go
#	docker run -it gqlgen-demo

generate:
	go generate ./...

