package controller

import (
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:    "jwt",
		Value:   "",
		Path:    "/",
		Expires: time.Now().AddDate(0, 0, -1),
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
