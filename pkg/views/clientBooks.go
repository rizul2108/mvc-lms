package views

import (
	"html/template"
)

func ClientBooksPage() *template.Template {

	temp := template.Must(template.ParseFiles("templates/booksClient.html"))
	return temp
}
