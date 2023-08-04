package views

import (
	"html/template"
)

func MakeAdminPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/makeAdmin.html"))
	return temp
}
