package controller

import (
	"mvc-go/pkg/models"
	"mvc-go/pkg/views"
	"net/http"
)

func AdminBooks(writer http.ResponseWriter, request *http.Request) {
	books, err := models.FetchBooks()
	if err != nil {
		http.Error(writer, "Database error", http.StatusInternalServerError)
		return
	}
	t := views.AdminBooksPage()
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, books)
}
