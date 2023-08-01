package views

import (
	"html/template"
)

func AddBookPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/addBook.html"))
	return temp
}
