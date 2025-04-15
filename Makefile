.PHONY:
all: build test run

run:
	./unit-converter

build:
	go build -o unit-converter .

test:
	go test ./...