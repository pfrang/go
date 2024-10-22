# Define variables
BINARY_NAME=out
BUILD_DIR=bin

# Default target: build the project
all: build

# Build the Go application
build:
	@echo "Building the application..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) .

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

# Help target to display available targets
help:
	@echo "Available targets:"
	@echo "  build   - Build the Go application"
	@echo "  clean   - Clean up build artifacts"
	@echo "  help    - Display this help message"

.PHONY: all build clean help
