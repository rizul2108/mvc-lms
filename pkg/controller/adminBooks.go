package controller

import (
	"mvc-go/pkg/models"
	"mvc-go/pkg/types"
	"mvc-go/pkg/views"
	"net/http"
)

func AdminBooks(writer http.ResponseWriter, _ *http.Request) {
	files := types.PutFileNames()

	books, err := models.FetchBooks()
	if err != nil {
		http.Error(writer, "Database error", http.StatusInternalServerError)
		return
	}
	t := views.ViewPage(files.AdminBooks)
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, books)
}

func AddQuantity(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookID")
	quantity := r.FormValue("quantity")
	err := models.AddQuantity(bookID, quantity)
	if err != "" {
		files := types.PutFileNames()

		t := views.ViewPage(files.AdminBooks)
		w.WriteHeader(http.StatusOK)
		t.Execute(w, err)
	} else {
		http.Redirect(w, r, "/admin/books", http.StatusSeeOther)
	}

}
func DecreaseQuantity(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookID")
	quantity := r.FormValue("quantity")
	err := models.DecreaseQuantity(bookID, quantity)
	if err != "" {
		files := types.PutFileNames()
		t := views.ViewPage(files.AdminBooks)
		w.WriteHeader(http.StatusOK)
		t.Execute(w, err)
	} else {
		http.Redirect(w, r, "/admin/books", http.StatusSeeOther)
	}

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookID")
	err := models.DeleteBook(bookID)
	if err != "" {
		files := types.PutFileNames()
		t := views.ViewPage(files.AdminBooks)
		w.WriteHeader(http.StatusOK)
		t.Execute(w, err)
	} else {
		http.Redirect(w, r, "/admin/books", http.StatusSeeOther)
	}
}
