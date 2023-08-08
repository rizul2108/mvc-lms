package controller

import (
	"fmt"
	"mvc-go/pkg/models"
	"mvc-go/pkg/types"
	"mvc-go/pkg/views"
	"net/http"
	"strconv"
)

func AdminRequests(w http.ResponseWriter, _ *http.Request) {
	files := types.PutFileNames()

	RequestList, error := models.FetchAllRequests()
	if error != "" {
		fmt.Println(error)
	} else {
		t := views.ViewPage(files.AdminRequests)
		w.WriteHeader(http.StatusOK)
		t.Execute(w, RequestList)
	}
}

func AcceptRequest(w http.ResponseWriter, r *http.Request) {
	requestId := r.FormValue("requestId")
	requestID, err := strconv.Atoi(requestId)
	if err != nil {

	}
	errorMessage := models.AcceptRequest(requestID)
	if errorMessage != "" {
		fmt.Println(errorMessage)
	}
	http.Redirect(w, r, "/admin/requests", http.StatusSeeOther)
}

func DeclineRequest(w http.ResponseWriter, r *http.Request) {
	requestId := r.FormValue("requestId")
	requestID, err := strconv.Atoi(requestId)
	if err != nil {

	}
	errorMessage := models.DeclineRequest(requestID)
	if errorMessage != "" {
		fmt.Println(errorMessage)
	}
	http.Redirect(w, r, "/admin/requests", http.StatusSeeOther)

}
