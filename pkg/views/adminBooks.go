package views

import (
	"html/template"
)

func AdminBooksPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/adminBooks.html"))
	return temp
}
