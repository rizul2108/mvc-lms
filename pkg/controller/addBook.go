package controller

import (
	"fmt"
	"mvc-go/pkg/models"
	"mvc-go/pkg/views"
	"net/http"
	"strconv"
)

func AddBook(w http.ResponseWriter, _ *http.Request) {
	t := views.ViewPage("addBook")
	w.WriteHeader(http.StatusOK)
	t.Execute(w, nil)
}

func AddNewBook(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	author := r.FormValue("author")
	quantityStr := r.FormValue("quantity")
	quantity, err := strconv.Atoi(quantityStr)
	fmt.Println(err)
	ErrorMessage := models.AddBook(title, author, quantity)
	if ErrorMessage != "" {
		t := views.ViewPage("makeAdmin")
		w.WriteHeader(http.StatusOK)
		t.Execute(w, ErrorMessage)
	} else {
		http.Redirect(w, r, "/admin/books", http.StatusSeeOther)
	}
}
