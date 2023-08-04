package models

import (
	"database/sql"
	"fmt"
	"mvc-go/pkg/types"
)

func FetchRequests(db *sql.DB, username string) ([]types.Request, string) {
	var userID int
	err := db.QueryRow("SELECT user_id FROM users WHERE username=?", username).Scan(&userID)
	if err != nil {
		fmt.Println(err)
		return nil, "User not found"
	}

	rows, err := db.Query("SELECT request_id,book_id, state, req_type FROM requests where user_id=?", userID)
	if err != nil {
		fmt.Println(err)
		return nil, "Internal Server Error 1"
	}
	defer rows.Close()
	var ownerName string
	db.QueryRow("SELECT full_name from users where user_id=?", userID).Scan(&ownerName)

	var requests []types.Request
	for rows.Next() {
		var request types.Request
		err := rows.Scan(&request.RequestID, &request.BookID, &request.State, &request.RequestType)
		if err != nil {
			fmt.Println(err)
			return nil, "Internal Server Error 2"
		}
		// Use a different variable name for the second result set
		var bookTitle string
		err = db.QueryRow("SELECT title from books where book_id=?", request.BookID).Scan(&bookTitle)
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
		SELECT r.request_id, r.book_id, r.state, r.req_type, b.title, u.full_name, r.user_id
		FROM requests r
		JOIN books b ON r.book_id = b.book_id
		JOIN users u ON r.user_id = u.user_id
		WHERE (r.req_type = 'return') OR (r.req_type = 'borrow' AND b.quantity > 0)`)
	if err != nil {
		fmt.Println(err)
		return nil, "Internal Server Error"
	}
	defer rows.Close()

	var requests []types.Request
	for rows.Next() {
		var request types.Request
		err := rows.Scan(&request.RequestID, &request.BookID, &request.State, &request.RequestType, &request.BookTitle, &request.OwnerName, &userID)
		if err != nil {
			fmt.Println(err)
			return nil, "Internal Server Error"
		}
		requests = append(requests, request)
	}

	return requests, ""
}
