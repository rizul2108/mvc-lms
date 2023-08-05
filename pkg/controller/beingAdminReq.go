package controller

import (
	"fmt"
	"mvc-go/pkg/models"
	"mvc-go/pkg/views"
	"net/http"
	"strconv"
)

func BeingAdminRequests(w http.ResponseWriter, r *http.Request) {
	db, err := models.Connection()
	if err != nil {
		return
	}
	AdminRequestList, errorMsg := models.FetchAdminRequests(db)
	if errorMsg != "" {
		fmt.Println(errorMsg)
	} else {
		t := views.BeingAdminRequestsPage()
		w.WriteHeader(http.StatusOK)
		t.Execute(w, AdminRequestList)
	}
}

func AcceptAdmin(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("userID")
	userId, err := strconv.Atoi(userID)
	if err != nil {

	}
	errorMessage := models.AcceptAdminReq(userId)
	if errorMessage != "" {
		fmt.Println(errorMessage)
	}
	http.Redirect(w, r, "/admin/adminRequests", http.StatusSeeOther)

}

func DeclineAdmin(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("userID")
	userId, err := strconv.Atoi(userID)
	if err != nil {

	}
	errorMessage := models.DeclineAdminReq(userId)
	if errorMessage != "" {
		fmt.Println(errorMessage)
	}
	http.Redirect(w, r, "/admin/adminRequests", http.StatusSeeOther)
}
