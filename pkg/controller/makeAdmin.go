package controller

import (
	"fmt"
	"mvc-go/pkg/models"
	"mvc-go/pkg/types"
	"mvc-go/pkg/views"
	"net/http"
)

func MakeAdmin(writer http.ResponseWriter, request *http.Request) {
	t := views.MakeAdminPage()
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}

func AddAdmin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("Username")
	password := r.FormValue("Password")
	passwordC := r.FormValue("PasswordConfirmVal")
	fullname := r.FormValue("Fullname")
	var errorMsg types.ErrorMessage
	str, errorMsg := models.AddUser(username, password, passwordC, fullname, "admin")
	if errorMsg.Message != "" {
		fmt.Print(errorMsg.Message)
		t := views.MakeAdminPage()
		w.WriteHeader(http.StatusOK)
		t.Execute(w, errorMsg)
	} else {
		fmt.Println(str)
		http.Redirect(w, r, "/admin/books", http.StatusSeeOther)
	}
}
