package views

import "mvc-go/pkg/types"

func PutFileNames() types.FileNames {
	return types.FileNames{
		AddBook:           "addBook.html",
		AdminBooks:        "adminBooks.html",
		AdminRequests:     "adminRequests.html",
		GetAdminRequests:  "getAdminRequests.html",
		BooksClient:       "booksClient.html",
		Profile:           "clientProfile.html",
		ErrorNotFound:     "errorNotFound.html",
		Home:              "home.html",
		Welcome:           "welcome.html",
		Login:             "login.html",
		Signup:            "signup.html",
		MakeAdmin:         "makeAdmin.html",
		AdminServerError:  "adminServerError.html",
		ClientServerError: "clientServerError.html",
	}
}
