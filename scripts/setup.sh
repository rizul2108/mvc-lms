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
if ! commandExists "python3"; then
  echo "Python is required not installed. Aborting."
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
read -p "Enter username of your first admin: " ADMIN_USERNAME
read -s -p "Enter password of your first admin: " ADMIN_PASSWORD
echo

migrate -path database/migration/ -database "mysql://$DB_USERNAME:$DB_PASSWORD@tcp(localhost:3306)/$DB_NAME" -verbose up

cat <<EOF > config.yaml
DB_USERNAME: $DB_USERNAME
DB_PASSWORD: '$DB_PASSWORD'
DB_HOST: $DB_HOST
DB_NAME: $DB_NAME
JWT_SECRET: "$JWT_SECRET"
EOF

echo "config.yaml created successfully with the provided values."

sudo apt install python3-pip
pip install mysql-connector-python

python3 <<EOF
import mysql.connector
import bcrypt

# Database connection parameters
db_username = "$DB_USERNAME"
db_password = "$DB_PASSWORD"
db_host = "$DB_HOST"
db_name = "$DB_NAME"

# User information
admin_username = "$ADMIN_USERNAME"
hashed_password = bcrypt.hashpw('$ADMIN_PASSWORD'.encode('utf-8'), bcrypt.gensalt()).decode('utf-8')

try:
    connection = mysql.connector.connect(user=db_username, password=db_password, host=db_host, database=db_name)
    cursor = connection.cursor()

    insert_query = "INSERT INTO users (username, fullName, hash, type) VALUES (%s, %s, %s, %s)"
    user_data = (admin_username, 'iamadmin', hashed_password, 'admin')

    cursor.execute(insert_query, user_data)
    connection.commit()

    print("Admin user inserted successfully!")

except mysql.connector.Error as error:
    print("Error:", error)

finally:
    if connection.is_connected():
        cursor.close()
        connection.close()
EOF

# Build the server binary
go mod vendor
go mod tidy
go build -o mvc ./cmd/main.go
if [ $? -ne 0 ]; then
  echo "Error building the server binary. Aborting."
  exit 1
fi

#For Virtual Hosting

read -p "Enter your domain name: " domain_name
read -p "Enter your email: " server_admin_email

sudo apt update
sudo apt install apache2

sudo a2enmod proxy proxy_http

sudo tee /etc/apache2/sites-available/$domain_name.conf > /dev/null <<EOF
<VirtualHost *:80>
    ServerName $domain_name
    ServerAdmin $server_admin_email
    ProxyPreserveHost On
    ProxyPass / http://127.0.0.1:9000/
    ProxyPassReverse / http://127.0.0.1:9000/
    TransferLog /var/log/apache2/${domain_name}_access.log
    ErrorLog /var/log/apache2/${domain_name}_error.log
</VirtualHost>
EOF

sudo a2ensite $domain_name.conf

sudo a2dissite 000-default.conf

sudo apache2ctl configtest

echo "127.0.0.1  $domain_name" | sudo tee -a /etc/hosts

sudo systemctl restart apache2


# Run the server
./mvc
if [ $? -ne 0 ]; then
  echo "Error running the server. Aborting."
  exit 1
fi

# Open the website in the browser
if commandExists "xdg-open"; then
  xdg-open http://$domain_name
elif commandExists "open"; then
  open http://$domain_name
else
  echo "Couldn't automatically open the website in the browser. You can manually open http://$domain_name"
fi
