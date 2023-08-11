package controller

import (
	"fmt"
	"mvc-go/pkg/models"
	"mvc-go/pkg/views"
	"net/http"
	"strconv"
)

func GetAdminRequests(w http.ResponseWriter, r *http.Request) {
	files := views.PutFileNames()

	db, err := models.Connection()
	if err != nil {
		return
	}
	adminRequestList, errorMsg := models.FetchAdminRequests(db)
	if errorMsg != "" {
		http.Redirect(w, r, "/admin/serverError", http.StatusSeeOther)
	} else {
		t := views.ViewPage(files.GetAdminRequests)
		w.WriteHeader(http.StatusOK)
		t.Execute(w, adminRequestList)
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
		http.Redirect(w, r, "/admin/serverError", http.StatusSeeOther)
	}
	http.Redirect(w, r, "/admin/adminRequests", http.StatusSeeOther)
}
