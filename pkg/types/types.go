package types

import "time"

type User struct {
	Fullname           string
	Username           string
	Password           string
	PasswordConfirmVal string
}

type Book struct {
	Quantity       int
	IssuedQuantity int
	Title          string
	Author         string
	ID             int
}
type Request struct {
	RequestID         int
	BookID            int
	State             string
	RequestType       string
	BookTitle         string
	OwnerName         string
	RequestDate       time.Time
	RequestDateString string
	Fine              int
	BookQuantity      int
}
type AdminRequest struct {
	UserID   int
	Username string
	Fullname string
}

type ErrorMessage struct {
	Message string `json:"message"`
}
