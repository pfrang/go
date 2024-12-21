# Variables
SRC_DIR := ./src
BIN_DIR := ./bin
MAIN_FILE := main.go
OUTPUT_BIN := main

# Default target
.PHONY: all
all: build

# Create bin directory if it doesn't exist
$(BIN_DIR):
	mkdir -p $(BIN_DIR)

# Build the Go binary
.PHONY: build
build: $(BIN_DIR)
	cd $(SRC_DIR) && go build -o ../$(BIN_DIR)/$(OUTPUT_BIN) $(MAIN_FILE)

# Clean up the bin directory
.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

# Run the binary
.PHONY: run
run: build
	./$(BIN_DIR)/$(OUTPUT_BIN)
