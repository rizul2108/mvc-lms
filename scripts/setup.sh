#!/bin/bash

# Function to check if a command is available
command_exists() {
  command -v "$1" >/dev/null 2>&1
}

# Check if 'mysql' command is available
if ! command_exists "mysql"; then
  echo "MySQL client is required but not installed. Aborting."
  exit 1
fi

# Check if 'go' command is available
if ! command_exists "go"; then
  echo "Go is required but not installed. Aborting."
  exit 1
fi

# Set up MySQL
read -p "Enter your MySQL root password: " mysql_password
mysql -u root -p"$mysql_password" < dump.sql
if [ $? -ne 0 ]; then
  echo "Error setting up MySQL. Aborting."
  exit 1
fi

read -p "Enter your database username: " DB_USERNAME
read -s -p "Enter your database password: " DB_PASSWORD
echo
read -p "Enter your database host: " DB_HOST
read -p "Enter your database name: " DB_NAME
read -p "Enter your JWT secret: " JWT_SECRET

cat <<EOF > config.yaml
DB_USERNAME: $DB_USERNAME
DB_PASSWORD: '$DB_PASSWORD'
DB_HOST: $DB_HOST
DB_NAME: $DB_NAME
JWT_SECRET: "$JWT_SECRET"
EOF

cp config.yaml ./pkg/models/config.yaml

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
if command_exists "xdg-open"; then
  xdg-open http://localhost:9000
elif command_exists "open"; then
  open http://localhost:9000
else
  echo "Couldn't automatically open the website in the browser. You can manually open http://localhost:9000"
fi
