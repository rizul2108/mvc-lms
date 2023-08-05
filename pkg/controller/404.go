package controller

import (
	"mvc-go/pkg/views"
	"net/http"
)

func NotFound(writer http.ResponseWriter, request *http.Request) {
	t := views.ViewPage("404")
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}
