package controller

import (
	"mvc-go/pkg/views"
	"net/http"
)

func AdminServerError(writer http.ResponseWriter, _ *http.Request) {
	files := views.PutFileNames()

	t := views.ViewPage(files.AdminServerError)
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}
func ClientServerError(writer http.ResponseWriter, _ *http.Request) {
	files := views.PutFileNames()

	t := views.ViewPage(files.ClientServerError)
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}
