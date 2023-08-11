package controller

import (
	"mvc-go/pkg/views"
	"net/http"
)

func NotFound(writer http.ResponseWriter, _ *http.Request) {
	files := views.PutFileNames()

	t := views.ViewPage(files.ErrorNotFound)
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}
