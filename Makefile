NAME := travis_demo

all: build

pre-build:
	git pull git@github.com:fefine/travis_demo.git

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/$(NAME) cmd/main.go

build-mac:
	go build -o build/$(NAME) cmd/main.go

run: pre-build
	make build
	./build/$(NAME)

test:
	go test

.PHONY: build all run