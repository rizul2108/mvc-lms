package controller

import (
	"mvc-go/pkg/views"
	"net/http"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	t := views.ViewPage("home")
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}
