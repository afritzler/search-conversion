GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
ARCH=amd64
BINARY_NAME=search-conversion
BINARY_UNIX=$(BINARY_NAME)_linux_$(ARCH)
BINARY_WIN=$(BINARY_NAME)_win_$(ARCH).exe
BINARY_DARWIN=$(BINARY_NAME)_darwin_$(ARCH)

all: test build
all-cross: build-linux build-win build-darwin
build:
		$(GOBUILD) -o $(BINARY_NAME) -v
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
		rm -f $(BINARY_WIN)
		rm -f $(BINARY_DARWIN)

lint:
		golangci-lint run ./...

run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)

# Cross compilation
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
build-win:
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_WIN) -v
build-darwin:
		CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_DARWIN) -v