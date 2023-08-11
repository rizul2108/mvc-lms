#!/bin/bash

read -p "Enter your database username: " DB_USERNAME
read -s -p "Enter your database password: " DB_PASSWORD
echo
read -p "Enter your database host: " DB_HOST
read -p "Enter your database name: " DB_NAME
read -p "Enter your JWT secret: " JWT_SECRET
read -p "Enter username of your first admin: " ADMIN_USERNAME
read -s -p "Enter password of your first admin: " ADMIN_PASSWORD
echo

# Hash the admin password using bcrypt
# HASHED_PASSWORD=$(python -c "import bcrypt; print(bcrypt.hashpw('$ADMIN_PASSWORD'.encode('utf-8'), bcrypt.gensalt()).decode('utf-8'))")

cat <<EOF > config.yaml
DB_USERNAME: $DB_USERNAME
DB_PASSWORD: '$DB_PASSWORD'
DB_HOST: $DB_HOST
DB_NAME: $DB_NAME
JWT_SECRET: "$JWT_SECRET"
EOF

echo "config.yaml created successfully with the provided values."

migrate -path database/migration/ -database "mysql://$DB_USERNAME:$DB_PASSWORD@tcp(localhost:3306)/$DB_NAME" -verbose up

# sudo apt install python3-pip
# pip install mysql-connector-python

# python3 <<EOF
# import mysql.connector

# # Database connection parameters
# db_username = "$DB_USERNAME"
# db_password = "$DB_PASSWORD"
# db_host = "$DB_HOST"
# db_name = "$DB_NAME"

# # User information
# admin_username = "$ADMIN_USERNAME"
# hashed_password = "$HASHED_PASSWORD"

# try:
#     connection = mysql.connector.connect(user=db_username, password=db_password, host=db_host, database=db_name)
#     cursor = connection.cursor()

#     insert_query = "INSERT INTO users (username, fullName, hash, type) VALUES (%s, %s, %s, %s)"
#     user_data = (admin_username, 'iamadmin', hashed_password, 'admin')

#     cursor.execute(insert_query, user_data)
#     connection.commit()

#     print("Admin user inserted successfully!")

# except mysql.connector.Error as error:
#     print("Error:", error)

# finally:
#     if connection.is_connected():
#         cursor.close()
#         connection.close()
# EOF
