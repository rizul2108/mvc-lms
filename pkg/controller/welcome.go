package controller

import (
	"mvc-go/pkg/types"
	"mvc-go/pkg/views"
	"net/http"
)

func Welcome(writer http.ResponseWriter, _ *http.Request) {
	files := types.PutFileNames()

	t := views.ViewPage(files.Welcome)
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, nil)
}
