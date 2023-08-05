# Makefile

# Variables
EXEC_FILE := mvc

.PHONY: all setup build run open

# Default target
all: setup replace build run open

# Set up MySQL
setup:
	mysql -u root -p < dump.sql

#Replace credentials in sampleConfig.yaml
replace:
	chmod +x ./scripts/credentialsReplace.sh
	./scripts/credentialsReplace.sh

# Build the server binary
build:
	go mod vendor
	go mod tidy
	go build -o $(EXEC_FILE) ./cmd/main.go

#Run the test function
test:
	go test ./pkg/models

# Run the server
run:
	./$(EXEC_FILE)

# Open the website in the browser
open:
	open http://localhost:9000
