package controller

import (
	// "encoding/json"

	"mvc-go/pkg/models"
	"mvc-go/pkg/views"
	"net/http"
)

func LogIn(writer http.ResponseWriter, request *http.Request) {
	t := views.LoginPage()
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("Username")
	password := r.FormValue("Password")
	jwtToken, userType, errorMessage := models.LoginUser(username, password)
	if errorMessage.Message != "" {
		t := views.LoginPage()
		w.WriteHeader(http.StatusOK)
		t.Execute(w, errorMessage)
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    jwtToken,
			Path:     "/",
			HttpOnly: true,
		})
		if userType == "client" {
			http.Redirect(w, r, "/profile", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/admin/books", http.StatusSeeOther)
		}
	}
}
