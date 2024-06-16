.PHONY: build build-run dc gen test run lint

build:
	go build -o ./build/converter ./cmd/converter/main.go

build-run: build
	./build/converter

gen:
	go generate ./...

test:
	go test -v -coverprofile cover.out ./... && go tool cover -html=cover.out

run:
	go run -race ./cmd/converter/main.go

lint:
	golangci-lint run