package controller

import (
	"mvc-go/pkg/views"
	"net/http"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	t := views.HomePage()
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}
