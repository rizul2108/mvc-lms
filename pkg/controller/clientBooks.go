package controller

import (
	"fmt"
	"mvc-go/pkg/models"
	"mvc-go/pkg/types"
	"mvc-go/pkg/views"
	"net/http"
	"strconv"
	"strings"
)

func ClientBooks(writer http.ResponseWriter, _ *http.Request) {
	files := types.PutFileNames()

	books, err := models.FetchBooks()
	if err != nil {
		http.Error(writer, "Database error", http.StatusInternalServerError)
		return
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
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	username := claims.Username
	error := models.AddRequest(ID, username)
	if error != "" {
		fmt.Println(error)
	}
	http.Redirect(w, r, "/client/books", http.StatusSeeOther)

}
