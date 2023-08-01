package controller

import (
	"mvc-go/pkg/views"
	"net/http"
)

func Welcome(writer http.ResponseWriter, request *http.Request) {
	t := views.WelcomePage()
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}
