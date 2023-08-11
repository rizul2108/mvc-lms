package models

import (
	"database/sql"
	"fmt"
	"mvc-go/pkg/types"
	"time"
)

func FetchRequests(db *sql.DB, username string) ([]types.Request, string) {
	var userID int
	err := db.QueryRow("SELECT userID FROM users WHERE username=?", username).Scan(&userID)
	if err != nil {
		return nil, "User not found"
	}

	rows, err := db.Query("SELECT requestID, bookID, state, requestType, requestDate FROM requests WHERE userID=?", userID)
	if err != nil {
		return nil, "Internal Server Error 1"
	}
	defer rows.Close()
	var ownerName string
	db.QueryRow("SELECT fullName FROM users WHERE userID=?", userID).Scan(&ownerName)

	var requests []types.Request
	for rows.Next() {
		var request types.Request
		err := rows.Scan(&request.RequestID, &request.BookID, &request.State, &request.RequestType, &request.RequestDateString)
		if err != nil {
			return nil, "Internal Server Error 2"
		}
		request.RequestDate, err = time.Parse("2006-01-02 15:04:05", request.RequestDateString)
		if err != nil {
			return nil, "Internal Server Error 3"
		}
		fine := 0
		if request.State == "Owned" && request.RequestType == "Accepted" {
			hoursDiff := time.Since(request.RequestDate).Hours()

			daysDiff := int(hoursDiff) / 24
			daysAfterGrace := daysDiff - 7

			if daysAfterGrace > 0 {
				fine = daysAfterGrace * 10
			}
		}
		request.Fine = fine
		var bookTitle string
		err = db.QueryRow("SELECT title FROM books WHERE bookID=?", request.BookID).Scan(&bookTitle)
		if err != nil {
			return nil, "Internal Server Error 3"
		}
		request.BookTitle = bookTitle
		request.OwnerName = ownerName
		requests = append(requests, request)
	}

	return requests, ""
}

func FetchAllRequests() ([]types.Request, string) {
	db, err := Connection()
	if err != nil {
		return nil, "Internal Server Error 2"
	}
	defer db.Close()

	var userID int
	rows, err := db.Query(`
		SELECT r.requestID, r.bookID, r.state, r.requestType, b.title, u.fullName, r.userID, r.requestDate
		FROM requests r
		JOIN books b ON r.bookID = b.bookID
		JOIN users u ON r.userID = u.userID
		WHERE (r.requestType = 'return') OR (r.requestType = 'Borrow')`)
	if err != nil {
		return nil, "Internal Server Error"
	}
	defer rows.Close()

	var requests []types.Request
	for rows.Next() {
		var request types.Request
		err := rows.Scan(&request.RequestID, &request.BookID, &request.State, &request.RequestType, &request.BookTitle, &request.OwnerName, &userID, &request.RequestDateString)
		if err != nil {
			return nil, "Internal Server Error"
		}
		request.RequestDate, err = time.Parse("2006-01-02 15:04:05", request.RequestDateString)
		fine := 0
		if request.State == "Owned" && request.RequestType == "Accepted" {
			hoursDiff := time.Since(request.RequestDate).Hours()

			daysDiff := int(hoursDiff) / 24
			daysAfterGrace := daysDiff - 7

			if daysAfterGrace > 0 {
				fine = daysAfterGrace * 10
			}
		}
		request.Fine = fine
		request.BookQuantity = 1
		if request.RequestType == "Borrow" {
			db.QueryRow(`SELECT quantity FROM books WHERE bookID =?`, request.BookID).Scan(&request.BookQuantity)
		}
		fmt.Println(request.BookQuantity)
		requests = append(requests, request)
	}

	return requests, ""
}

func FetchAdminRequests(db *sql.DB) ([]types.AdminRequest, string) {
	rows, err := db.Query(`SELECT userID,username, fullName FROM users WHERE type="Requested"`)
	if err != nil {
		return nil, "Internal Server Error 1"
	}
	defer rows.Close()

	var adminRequests []types.AdminRequest
	for rows.Next() {
		var adminRequest types.AdminRequest
		err := rows.Scan(&adminRequest.UserID, &adminRequest.Username, &adminRequest.Fullname)
		if err != nil {
			return nil, "Internal Server Error 2"
		}
		adminRequests = append(adminRequests, adminRequest)
	}

	return adminRequests, ""
}
