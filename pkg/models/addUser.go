package models

import (
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func AddUser(username, password, passwordC, fullname, user_type string) (string, string) {
	if password != passwordC {
		return "", "Passwords didn't match"
	}

	db, err := Connection()
	if err != nil {
		return "", "Internal Server Error 1"
	}
	defer db.Close()

	var userExists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE username = ?)", username).Scan(&userExists)
	if err != nil {
		return "", "Internal Server Error 2"
	}

	if userExists {
		return "", "Username Already Exists "
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return "", "Internal Server Error 3"
	}

	_, err = db.Exec(`INSERT INTO users (username, full_name,hash,type) VALUES (?, ?, ?,?)`, username, fullname, hashedPassword, user_type)
	if err != nil {
		return "", "Internal Server Error 4"
	}

	jwt, err := GenerateToken(username)
	if err != nil {
		return "", "Error in producing token"
	}
	return jwt, ""

}
