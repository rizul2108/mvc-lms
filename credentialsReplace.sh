#!/bin/bash

# Prompt user for database credentials
read -p "Enter your database username: " DB_USERNAME
read -s -p "Enter your database password: " DB_PASSWORD
echo
read -p "Enter your database host: " DB_HOST
read -p "Enter your database name: " DB_NAME

# Replace credentials in sampleConfig.yaml
sed -i "s/DB_USERNAME_PLACEHOLDER/$DB_USERNAME/g" sampleConfig.yaml
sed -i "s/DB_PASSWORD_PLACEHOLDER/$DB_PASSWORD/g" sampleConfig.yaml
sed -i "s/DB_HOST_PLACEHOLDER/$DB_HOST/g" sampleConfig.yaml
sed -i "s/DB_NAME_PLACEHOLDER/$DB_NAME/g" sampleConfig.yaml

# Rename sampleConfig.yaml to config.yaml
mv sampleConfig.yaml config.yaml

# Repeat the same process for ./pkg/models/sampleConfig.yaml
sed -i "s/DB_USERNAME_PLACEHOLDER/$DB_USERNAME/g" ./pkg/models/sampleConfig.yaml
sed -i "s/DB_PASSWORD_PLACEHOLDER/$DB_PASSWORD/g" ./pkg/models/sampleConfig.yaml
sed -i "s/DB_HOST_PLACEHOLDER/$DB_HOST/g" ./pkg/models/sampleConfig.yaml
sed -i "s/DB_NAME_PLACEHOLDER/$DB_NAME/g" ./pkg/models/sampleConfig.yaml
mv ./pkg/models/sampleConfig.yaml ./pkg/models/config.yaml
