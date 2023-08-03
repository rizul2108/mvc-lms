package controller

import (
	"fmt"
	"mvc-go/pkg/models"
	"mvc-go/pkg/types"
	"mvc-go/pkg/views"
	"net/http"
)

func SignUp(writer http.ResponseWriter, request *http.Request) {
	t := views.SignUpPage()
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}
func AddUser(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("Username")
	password := r.FormValue("Password")
	passwordC := r.FormValue("PasswordConfirmVal")
	fullname := r.FormValue("Fullname")
	var errorMessage types.ErrorMessage
	var str string
	str, errorMessage = models.AddUser(username, password, passwordC, fullname, "client")
	if errorMessage.Message != "" {
		fmt.Println(errorMessage.Message)
		t := views.SignUpPage()
		w.WriteHeader(http.StatusOK)
		t.Execute(w, errorMessage)
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    str,
			Path:     "/",
			HttpOnly: true,
		})
		http.Redirect(w, r, "/client/profile", http.StatusSeeOther)
	}
}
