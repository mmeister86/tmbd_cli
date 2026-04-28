BINARY_NAME=tmdb
VERSION=1.0.2
VERSION_LDFLAGS=-X github.com/mmeister86/tmbd_cli/cmd.Version=$(VERSION)
INSTALL_DIR=/usr/local/bin
SUDO=sudo

.PHONY: all build clean install test build-all

all: build

build:
	go build -ldflags "$(VERSION_LDFLAGS)" -o $(BINARY_NAME) .

build-all:
	@mkdir -p dist
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(VERSION_LDFLAGS)" -o dist/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -ldflags "$(VERSION_LDFLAGS)" -o dist/$(BINARY_NAME)-darwin-arm64 .
	GOOS=linux GOARCH=amd64 go build -ldflags "$(VERSION_LDFLAGS)" -o dist/$(BINARY_NAME)-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -ldflags "$(VERSION_LDFLAGS)" -o dist/$(BINARY_NAME)-linux-arm64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "$(VERSION_LDFLAGS)" -o dist/$(BINARY_NAME)-windows-amd64.exe .

install: build
	$(SUDO) mkdir -p $(INSTALL_DIR)
	$(SUDO) mv $(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)

uninstall:
	$(SUDO) rm -f $(INSTALL_DIR)/$(BINARY_NAME)

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
