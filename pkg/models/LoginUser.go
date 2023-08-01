package models

import (
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(username, password string) (string, string, string) {
	db, err := Connection()
	if err != nil {
		return "", "", "Internal Server Error 1"
	}
	defer db.Close()

	var userExists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE username = ?)", username).Scan(&userExists)
	if err != nil {
		return "", "", "Internal Server Error 2"
	}

	if !userExists {
		return "", "", "Username Doesn't Exist"
	}

	var hashedPassword, userType string
	err = db.QueryRow("SELECT hash,type FROM users WHERE username = ?", username).Scan(&hashedPassword, &userType)
	if err != nil {
		return "", "", "Internal Server Error 3"
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return "", "", "Invalid Password"
	}

	jwtToken, err := generateToken(username)
	if err != nil {
		return "", "", "Error in producing token"
	}

	return jwtToken, userType, ""
}
