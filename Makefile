# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary name
BINARY_NAME=url-shortener

# Main build target
all: test build

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Run tests
test:
	$(GOTEST) -v ./...

# Clean up
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Install dependencies
deps:
	$(GOGET) -u github.com/gofiber/fiber/v2
	$(GOGET) -u gorm.io/gorm/logger
	$(GOGET) -u github.com/joho/godotenv
	$(GOGET) -u github.com/go-redis/redis/v8
	$(GOGET) -u github.com/asaskevich/govalidator
	$(GOGET) -u github.com/go-delve/delve/cmd/dlv

# Run the application
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

#background dependencies
tidy:
	$(GOCMD) mod tidy

# Shortcuts
.PHONY: all build test clean deps run tidy
