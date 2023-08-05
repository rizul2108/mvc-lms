package controller

import (
	"mvc-go/pkg/views"
	"net/http"
)

func Welcome(writer http.ResponseWriter, _ *http.Request) {
	t := views.ViewPage("welcome")
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}
