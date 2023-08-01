package views

import (
	"html/template"
)

func SignUpPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/signup.html"))
	return temp
}
