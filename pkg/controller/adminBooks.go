package controller

import (
	"mvc-go/pkg/models"
	"mvc-go/pkg/views"
	"net/http"
)

func AdminBooks(writer http.ResponseWriter, r *http.Request) {
	files := views.PutFileNames()

	books, err := models.FetchBooks()
	if err != "" {
		http.Redirect(writer, r, "/admin/serverError", http.StatusSeeOther)
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
		http.Redirect(w, r, "/admin/serverError", http.StatusSeeOther)
		return
	} else {
		http.Redirect(w, r, "/admin/books", http.StatusSeeOther)
		return
	}

}
func DecreaseQuantity(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookID")
	quantity := r.FormValue("quantity")
	err := models.DecreaseQuantity(bookID, quantity)
	if err != "" {
		http.Redirect(w, r, "/admin/serverError", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/admin/books", http.StatusSeeOther)
		return
	}

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookID")
	err := models.DeleteBook(bookID)
	if err != "" {
		http.Redirect(w, r, "/admin/serverError", http.StatusSeeOther)
		return
	} else {
		http.Redirect(w, r, "/admin/books", http.StatusSeeOther)
		return
	}
}
