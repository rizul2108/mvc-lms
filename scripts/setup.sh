#!/bin/bash

# Function to check if a command is available
commandExists() {
  command -v "$1" >/dev/null 2>&1
}

# Check if 'mysql' command is available
if ! commandExists "mysql"; then
  echo "MySQL client is required but not installed. Aborting."
  exit 1
fi

# Check if 'go' command is available
if ! commandExists "go"; then
  echo "Go is required but not installed. Aborting."
  exit 1
fi

read -p "Enter your database username: " DB_USERNAME
read -s -p "Enter your database password: " DB_PASSWORD
echo
read -p "Enter your database host: " DB_HOST
read -p "Enter your database name: " DB_NAME
read -p "Enter your JWT secret: " JWT_SECRET

migrate -path database/migration/ -database "mysql://$DB_USERNAME:$DB_PASSWORD@tcp(localhost:3306)/$DB_NAME" -verbose up

cat <<EOF > config.yaml
DB_USERNAME: $DB_USERNAME
DB_PASSWORD: '$DB_PASSWORD'
DB_HOST: $DB_HOST
DB_NAME: $DB_NAME
JWT_SECRET: "$JWT_SECRET"
EOF

echo "config.yaml created successfully with the provided values."

# Build the server binary
go mod vendor
go mod tidy
go build -o mvc ./cmd/main.go
if [ $? -ne 0 ]; then
  echo "Error building the server binary. Aborting."
  exit 1
fi

# Run the server
./mvc
if [ $? -ne 0 ]; then
  echo "Error running the server. Aborting."
  exit 1
fi

# Open the website in the browser
if commandExists "xdg-open"; then
  xdg-open http://localhost:9000
elif commandExists "open"; then
  open http://localhost:9000
else
  echo "Couldn't automatically open the website in the browser. You can manually open http://localhost:9000"
fi
