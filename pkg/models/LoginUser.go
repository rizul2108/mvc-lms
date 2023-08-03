package models

import (
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"mvc-go/pkg/types"
)

func LoginUser(username, password string) (string, string, types.ErrorMessage) {
	db, err := Connection()
	var errorMsg types.ErrorMessage
	if err != nil {
		errorMsg.Message = "Error in connecting to database"
		return "", "", errorMsg
	}
	defer db.Close()

	var userExists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE username = ?)", username).Scan(&userExists)
	if err != nil {
		errorMsg.Message = "Internal Server Error 2"
		return "", "", errorMsg
	}

	if !userExists {
		errorMsg.Message = "Username already exists"
		return "", "", errorMsg
	}

	var hashedPassword, userType string
	err = db.QueryRow("SELECT hash,type FROM users WHERE username = ?", username).Scan(&hashedPassword, &userType)
	if err != nil {
		errorMsg.Message = "Internal Server Error 3"
		return "", "", errorMsg
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		errorMsg.Message = "wrong Password"
		return "", "", errorMsg
	}

	jwtToken, err := GenerateToken(username)
	if err != nil {
		errorMsg.Message = "Internal server error in jwt"
		return "", "", errorMsg
	}
	errorMsg.Message = ""
	return jwtToken, userType, errorMsg
}
