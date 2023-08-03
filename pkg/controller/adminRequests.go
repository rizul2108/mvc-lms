package controller

import (
	"fmt"
	"mvc-go/pkg/models"
	"mvc-go/pkg/views"
	"net/http"
	"strconv"
)

func AdminRequests(w http.ResponseWriter, r *http.Request) {

	RequestList, error := models.FetchAllRequests()
	if error != "" {
		fmt.Println(error)
	} else {
		t := views.AdminRequestsPage()
		w.WriteHeader(http.StatusOK)
		t.Execute(w, RequestList)
	}
}

func AcceptRequest(w http.ResponseWriter, r *http.Request) {
	reqID := r.FormValue("reqID")
	requestID, err := strconv.Atoi(reqID)
	if err != nil {

	}
	errorMessage := models.AcceptRequest(requestID)
	if errorMessage != "" {
		fmt.Println(errorMessage)
	}
	http.Redirect(w, r, "/admin/requests", http.StatusSeeOther)
}

func DeclineRequest(w http.ResponseWriter, r *http.Request) {
	reqID := r.FormValue("reqID")
	requestID, err := strconv.Atoi(reqID)
	if err != nil {

	}
	errorMessage := models.DeclineRequest(requestID)
	if errorMessage != "" {
		fmt.Println(errorMessage)
	}
	http.Redirect(w, r, "/admin/requests", http.StatusSeeOther)

}
