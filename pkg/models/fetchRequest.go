package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mvc-go/pkg/types"
)

func FetchRequests(username string) ([]types.Request, string) {
	db, err := Connection()

	var userID int
	err = db.QueryRow("SELECT user_id FROM users WHERE username=?", username).Scan(&userID)
	if err != nil {
		return nil, "User not found"
	}

	rows, err := db.Query("SELECT book_id, state, req_type FROM requests where user_id=?", userID)
	if err != nil {
		fmt.Println(err)
		return nil, "Internal Server Error"
	}
	defer rows.Close()

	var requests []types.Request
	for rows.Next() {
		var request types.Request
		err := rows.Scan(&request.BookID, &request.State, &request.ReqType)
		if err != nil {
			fmt.Println(err)
			return nil, "Internal Server Error"
		}
		requests = append(requests, request)
	}

	return requests, ""
}
