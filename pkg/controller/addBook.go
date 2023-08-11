package controller

import (
	"fmt"
	"mvc-go/pkg/models"
	"mvc-go/pkg/views"
	"net/http"
	"strconv"
)

func AddBook(w http.ResponseWriter, _ *http.Request) {
	files := views.PutFileNames()

	t := views.ViewPage(files.AddBook)
	w.WriteHeader(http.StatusOK)
	t.Execute(w, nil)
}

func AddNewBook(w http.ResponseWriter, r *http.Request) {
	files := views.PutFileNames()

	title := r.FormValue("title")
	author := r.FormValue("author")
	quantityStr := r.FormValue("quantity")
	quantity, err := strconv.Atoi(quantityStr)
	fmt.Println(err)
	errorMessage := models.AddBook(title, author, quantity)
	if errorMessage != "" {
		t := views.ViewPage(files.MakeAdmin)
		w.WriteHeader(http.StatusOK)
		t.Execute(w, errorMessage)
	} else {
		http.Redirect(w, r, "/admin/books", http.StatusSeeOther)
	}
}
