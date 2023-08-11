EXEC_FILE := mvc

.PHONY: all replace&setupMySQL build test run open

all: replace&setupMySQL build test run open

# Replace credentials in sampleConfig.yaml
replace&setupMySQL:
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

