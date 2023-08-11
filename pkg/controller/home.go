package controller

import (
	"mvc-go/pkg/views"
	"net/http"
)

func Home(writer http.ResponseWriter, _ *http.Request) {
	files := views.PutFileNames()

	t := views.ViewPage(files.Home)
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}
