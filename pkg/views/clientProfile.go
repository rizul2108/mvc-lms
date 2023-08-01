package views

import (
	"html/template"
)

func ProfilePage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/clientProfile.html"))
	return temp
}
