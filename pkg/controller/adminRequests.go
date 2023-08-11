package controller

import (
	"mvc-go/pkg/models"
	"mvc-go/pkg/views"
	"net/http"
	"strconv"
)

func AdminRequests(w http.ResponseWriter, r *http.Request) {
	files := views.PutFileNames()

	requestList, error := models.FetchAllRequests()
	if error != "" {
		http.Redirect(w, r, "/admin/serverError", http.StatusSeeOther)
	} else {
		t := views.ViewPage(files.AdminRequests)
		w.WriteHeader(http.StatusOK)
		t.Execute(w, requestList)
	}
}

func AcceptRequest(w http.ResponseWriter, r *http.Request) {
	requestId := r.FormValue("requestId")
	requestID, err := strconv.Atoi(requestId)
	if err != nil {

	}
	errorMessage := models.AcceptRequest(requestID)
	if errorMessage != "" {
		http.Redirect(w, r, "/admin/serverError", http.StatusSeeOther)
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
		http.Redirect(w, r, "/admin/serverError", http.StatusSeeOther)
	}
	http.Redirect(w, r, "/admin/requests", http.StatusSeeOther)

}
