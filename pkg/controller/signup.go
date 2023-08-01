package controller

import (
	"fmt"
	"mvc-go/pkg/models"
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
	passwordC := r.FormValue("PasswordC")
	fullname := r.FormValue("Fullname")
	str, ErrorMessage := models.AddUser(username, password, passwordC, fullname, "client")
	if ErrorMessage != "" {
		fmt.Println(ErrorMessage)
		t := views.SignUpPage()
		w.WriteHeader(http.StatusOK)
		t.Execute(w, ErrorMessage)
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    str,
			Path:     "/",
			HttpOnly: true,
		})
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
	}
}
