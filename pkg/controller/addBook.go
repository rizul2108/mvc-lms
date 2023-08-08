package controller

import (
	"fmt"
	"mvc-go/pkg/models"
	"mvc-go/pkg/types"
	"mvc-go/pkg/views"
	"net/http"
	"strconv"
)

func AddBook(w http.ResponseWriter, _ *http.Request) {
	files := types.PutFileNames()

	t := views.ViewPage(files.AddBook)
	w.WriteHeader(http.StatusOK)
	t.Execute(w, nil)
}

func AddNewBook(w http.ResponseWriter, r *http.Request) {
	files := types.PutFileNames()

	title := r.FormValue("title")
	author := r.FormValue("author")
	quantityStr := r.FormValue("quantity")
	quantity, err := strconv.Atoi(quantityStr)
	fmt.Println(err)
	ErrorMessage := models.AddBook(title, author, quantity)
	if ErrorMessage != "" {
		t := views.ViewPage(files.MakeAdmin)
		w.WriteHeader(http.StatusOK)
		t.Execute(w, ErrorMessage)
	} else {
		http.Redirect(w, r, "/admin/books", http.StatusSeeOther)
	}
}
