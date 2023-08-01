package controller

import (
	"mvc-go/pkg/views"
	"net/http"
)

func Profile(writer http.ResponseWriter, request *http.Request) {
	t := views.ProfilePage()
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}
