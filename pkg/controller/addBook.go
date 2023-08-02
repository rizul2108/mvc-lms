package controller

import (
	"fmt"
	"mvc-go/pkg/models"
	"mvc-go/pkg/views"
	"net/http"
	"strconv"
)

func AddBook(writer http.ResponseWriter, request *http.Request) {
	t := views.AddBookPage()
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}
func AddNewBook(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	author := r.FormValue("author")
	qtyStr := r.FormValue("quantity")
	qty, err := strconv.Atoi(qtyStr)
	fmt.Println(err)
	ErrorMessage := models.AddBook(title, author, qty)
	if ErrorMessage != "" {
		t := views.MakeAdminPage()
		w.WriteHeader(http.StatusOK)
		t.Execute(w, ErrorMessage)
	} else {
		http.Redirect(w, r, "/adminBooks", http.StatusSeeOther)
	}
}
