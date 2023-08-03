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
func AddQuantity(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookID")
	quantity := r.FormValue("quantity")
	err := models.AddQuantity(bookID, quantity)
	if err != "" {
		t := views.MakeAdminPage()
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
		t := views.MakeAdminPage()
		w.WriteHeader(http.StatusOK)
		t.Execute(w, err)
	} else {
		http.Redirect(w, r, "/admin/books", http.StatusSeeOther)
	}

}
