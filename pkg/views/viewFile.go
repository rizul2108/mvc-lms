package views

import (
	"html/template"
)

func ViewPage(fileName string) *template.Template {
	temp := template.Must(template.ParseFiles("templates/" + fileName))
	return temp
}
