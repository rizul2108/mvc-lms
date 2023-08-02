package controller

import (
	"fmt"
	"mvc-go/pkg/models"
	"mvc-go/pkg/views"
	"net/http"
	"strconv"
	"strings"
)

func ClientBooks(writer http.ResponseWriter, request *http.Request) {
	books, err := models.FetchBooks()
	if err != nil {
		http.Error(writer, "Database error", http.StatusInternalServerError)
		return
	}
	t := views.ClientBooksPage()
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, books)
}

func AddRequest(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookID")
	cookie, err := r.Cookie("jwt")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	ID, err := strconv.Atoi(bookID)
	fmt.Println(bookID)
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
