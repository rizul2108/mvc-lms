package models

import (
	"golang.org/x/crypto/bcrypt"
	"mvc-go/pkg/types"
)

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func AddUser(username, password, passwordC, fullname, user_type string) (string, types.ErrorMessage) {
	var errorMsg types.ErrorMessage
	var jwt string = ""
	if password != passwordC {
		errorMsg.Message = "Passwords didn't match"
		return "", errorMsg
	}

	db, err := Connection()
	if err != nil {
		errorMsg.Message = "Passwords didn't match"
		return "", errorMsg
	}
	defer db.Close()

	var userExists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE username = ?)", username).Scan(&userExists)
	if err != nil {
		return "", errorMsg
	}

	if userExists {
		errorMsg.Message = "Username Already Exists "

		return "", errorMsg
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		errorMsg.Message = "Username Already Exists "
		return "", errorMsg
	}

	_, err = db.Exec(`INSERT INTO users (username, full_name,hash,type) VALUES (?, ?, ?,?)`, username, fullname, hashedPassword, user_type)
	if err != nil {
		errorMsg.Message = "Internal Server Error 4"
		return "", errorMsg
	}

	jwt, err = GenerateToken(username)
	if err != nil {
		errorMsg.Message = "Error in getting token"
		return "", errorMsg
	}
	errorMsg.Message = ""

	return jwt, errorMsg

}
