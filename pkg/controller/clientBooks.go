package controller

import (
	"mvc-go/pkg/models"
	"mvc-go/pkg/views"
	"net/http"
	"strconv"
	"strings"
)

func ClientBooks(writer http.ResponseWriter, request *http.Request) {
	files := views.PutFileNames()

	books, err := models.FetchBooks()
	if err != "" {
		http.Redirect(writer, request, "/client/serverError", http.StatusSeeOther)
	}
	t := views.ViewPage(files.BooksClient)
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, books)
}

func AddRequest(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookID")
	cookie, err := r.Cookie("jwt")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	ID, err := strconv.Atoi(bookID)
	tokenString := strings.TrimSpace(cookie.Value)
	claims, err := models.VerifyToken(tokenString)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	username := claims.Username
	error := models.AddRequest(ID, username)
	if error != "" {
		http.Redirect(w, r, "/client/serverError", http.StatusSeeOther)
	}
	http.Redirect(w, r, "/client/books", http.StatusSeeOther)

}
