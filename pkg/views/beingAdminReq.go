package views

import (
	"html/template"
)

func BeingAdminRequestsPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/beingAdminReq.html"))
	return temp
}
