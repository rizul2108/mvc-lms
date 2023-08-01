package views

import (
	"html/template"
)

func WelcomePage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/welcome.html"))
	return temp
}
