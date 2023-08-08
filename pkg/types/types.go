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

type FileNames struct {
	AddBook            string
	AdminBooks         string
	AdminRequests      string
	BeingAdminRequests string
	BooksClient        string
	Profile            string
	ErrorNotFound      string
	Home               string
	Welcome            string
	Login              string
	Signup             string
	MakeAdmin          string
}

func PutFileNames() FileNames {
	return FileNames{
		AddBook:            "addBook.html",
		AdminBooks:         "adminBooks.html",
		AdminRequests:      "adminRequests.html",
		BeingAdminRequests: "beingAdminReq.html",
		BooksClient:        "booksClient.html",
		Profile:            "clientProfile.html",
		ErrorNotFound:      "errorNotFound.html",
		Home:               "home.html",
		Welcome:            "welcome.html",
		Login:              "login.html",
		Signup:             "signup.html",
		MakeAdmin:          "makeAdmin.html",
	}
}
