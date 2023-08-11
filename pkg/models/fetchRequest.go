package models

import (
	"database/sql"
	"fmt"
	"mvc-go/pkg/types"
	"time"
)

func FetchRequests(db *sql.DB, username string) ([]types.Request, string) {
	var userID int
	err := db.QueryRow("SELECT user_id FROM users WHERE username=?", username).Scan(&userID)
	if err != nil {
		fmt.Println(err)
		return nil, "User not found"
	}

	rows, err := db.Query("SELECT requestID, book_id, state, requestType, requestDate FROM requests WHERE user_id=?", userID)
	if err != nil {
		fmt.Println(err)
		return nil, "Internal Server Error 1"
	}
	defer rows.Close()
	var ownerName string
	db.QueryRow("SELECT full_name FROM users WHERE user_id=?", userID).Scan(&ownerName)

	var requests []types.Request
	for rows.Next() {
		var request types.Request
		err := rows.Scan(&request.RequestID, &request.BookID, &request.State, &request.RequestType, &request.RequestDateString)
		if err != nil {
			fmt.Println(err)
			return nil, "Internal Server Error 2"
		}
		request.RequestDate, err = time.Parse("2006-01-02 15:04:05", request.RequestDateString)
		if err != nil {
			fmt.Println(err)
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
		err = db.QueryRow("SELECT title FROM books WHERE book_id=?", request.BookID).Scan(&bookTitle)
		if err != nil {
			fmt.Println(err)
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
		fmt.Println(err)
		return nil, "Internal Server Error 2"
	}
	defer db.Close()

	var userID int
	rows, err := db.Query(`
		SELECT r.requestID, r.book_id, r.state, r.requestType, b.title, u.full_name, r.user_id, r.requestDate
		FROM requests r
		JOIN books b ON r.book_id = b.book_id
		JOIN users u ON r.user_id = u.user_id
		WHERE (r.requestType = 'return') OR (r.requestType = 'Borrow')`)
	if err != nil {
		fmt.Println(err)
		return nil, "Internal Server Error"
	}
	defer rows.Close()

	var requests []types.Request
	for rows.Next() {
		var request types.Request
		err := rows.Scan(&request.RequestID, &request.BookID, &request.State, &request.RequestType, &request.BookTitle, &request.OwnerName, &userID, &request.RequestDateString)
		if err != nil {
			fmt.Println(err)
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
		db.QueryRow(`SELECT COUNT(*) FROM requests WHERE book_id =?`, request.BookID).Scan(&request.BookQuantity)
		requests = append(requests, request)
	}

	return requests, ""
}

func FetchAdminRequests(db *sql.DB) ([]types.AdminRequest, string) {
	rows, err := db.Query(`SELECT user_id,username, full_name FROM users WHERE type="Requested"`)
	if err != nil {
		fmt.Println(err)
		return nil, "Internal Server Error 1"
	}
	defer rows.Close()

	var adminRequests []types.AdminRequest
	for rows.Next() {
		var adminRequest types.AdminRequest
		err := rows.Scan(&adminRequest.UserID, &adminRequest.Username, &adminRequest.Fullname)
		if err != nil {
			fmt.Println(err)
			return nil, "Internal Server Error 2"
		}
		adminRequests = append(adminRequests, adminRequest)
	}

	return adminRequests, ""
}
