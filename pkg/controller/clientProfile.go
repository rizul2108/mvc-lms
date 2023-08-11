package controller

import (
	"fmt"
	"mvc-go/pkg/models"
	"mvc-go/pkg/types"
	"mvc-go/pkg/views"
	"net/http"
	"strconv"
	"strings"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	files := types.PutFileNames()

	cookie, err := r.Cookie("jwt")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	tokenString := strings.TrimSpace(cookie.Value)
	claims, err := models.VerifyToken(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	db, err := models.Connection()
	username := claims.Username

	ReqList, error := models.FetchRequests(db, username)
	if error != "" {
		fmt.Println(error)
	} else {
		t := views.ViewPage(files.Profile)
		w.WriteHeader(http.StatusOK)
		t.Execute(w, ReqList)
	}

}

func DeleteRequest(w http.ResponseWriter, r *http.Request) {
	requestID := r.FormValue("requestId")
	requestId, error := strconv.Atoi(requestID)
	if error == nil {
		err := models.DeleteRequest(requestId)
		if err == "" {
			http.Redirect(w, r, "/client/profile", http.StatusSeeOther)
		}
	}
}

func ReturnBook(w http.ResponseWriter, r *http.Request) {
	requestID := r.FormValue("requestId")
	requestId, error := strconv.Atoi(requestID)
	fmt.Println(error)
	if error == nil {
		err := models.ReturnBook(requestId)
		if err == "" {
			http.Redirect(w, r, "/client/profile", http.StatusSeeOther)
		} else {

		}
	}
}
