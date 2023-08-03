package views

import (
	"html/template"
)

func MakeAdminPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/admin/makeAdmin.html"))
	return temp
}
