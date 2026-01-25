BINARY_NAME=tmdb
VERSION=1.0.0

.PHONY: all build clean install test build-all

all: build

build:
	go build -ldflags "-X main.Version=$(VERSION)" -o $(BINARY_NAME) .

build-all:
	@mkdir -p dist
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o dist/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -ldflags "-X main.Version=$(VERSION)" -o dist/$(BINARY_NAME)-darwin-arm64 .
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o dist/$(BINARY_NAME)-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -ldflags "-X main.Version=$(VERSION)" -o dist/$(BINARY_NAME)-linux-arm64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o dist/$(BINARY_NAME)-windows-amd64.exe .

install: build
	sudo mv $(BINARY_NAME) /usr/local/bin/

uninstall:
	sudo rm -f /usr/local/bin/$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
	rm -rf dist/

test:
	go test -v ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

lint: fmt vet
	@echo "Linting complete"
