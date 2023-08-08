#!/bin/bash

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

echo "config.yaml created successfully with the provided values."

migrate -path database/migration/ -database "mysql://$DB_USERNAME:$DB_PASSWORD@tcp(localhost:3306)/$DB_NAME" -verbose up
