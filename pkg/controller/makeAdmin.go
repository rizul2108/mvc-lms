package controller

import (
	"fmt"
	"mvc-go/pkg/models"
	"mvc-go/pkg/types"
	"mvc-go/pkg/views"
	"net/http"
)

func MakeAdmin(writer http.ResponseWriter, _ *http.Request) {
	files := views.PutFileNames()

	t := views.ViewPage(files.MakeAdmin)
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}

func AddAdmin(w http.ResponseWriter, r *http.Request) {
	files := views.PutFileNames()

	username := r.FormValue("Username")
	password := r.FormValue("Password")
	passwordC := r.FormValue("PasswordConfirmVal")
	fullname := r.FormValue("Fullname")
	var errorMsg types.ErrorMessage
	str, errorMsg := models.AddUser(username, password, passwordC, fullname, "admin")
	if errorMsg.Message != "" {
		fmt.Print(errorMsg.Message)
		t := views.ViewPage(files.MakeAdmin)
		t.Execute(w, errorMsg)
	} else {
		fmt.Println(str)
		http.Redirect(w, r, "/admin/books", http.StatusSeeOther)
	}
}
