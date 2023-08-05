package controller

import (
	"fmt"
	"mvc-go/pkg/models"
	"mvc-go/pkg/types"
	"mvc-go/pkg/views"
	"net/http"
)

func SignUp(writer http.ResponseWriter, _ *http.Request) {
	t := views.ViewPage("signup")
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("Username")
	password := r.FormValue("Password")
	passwordC := r.FormValue("PasswordConfirmVal")
	fullname := r.FormValue("Fullname")

	adminRequest := r.Form["adminRequest"] != nil
	fmt.Println(r.Form["adminRequest"])

	var errorMessage types.ErrorMessage
	var str string
	fmt.Println(adminRequest)
	if adminRequest == false {
		str, errorMessage = models.AddUser(username, password, passwordC, fullname, "client")
	} else {
		str, errorMessage = models.AddUser(username, password, passwordC, fullname, "requested")
	}

	if errorMessage.Message != "" {
		fmt.Println(errorMessage.Message)
		t := views.ViewPage("signup")
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
