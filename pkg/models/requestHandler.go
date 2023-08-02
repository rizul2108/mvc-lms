package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func AddRequest(bookID int, username string) string {
	db, err := Connection()
	if err != nil {
		return "Internal Server Error 1"
	}

	var userID int
	err = db.QueryRow("SELECT user_id FROM users WHERE username=?", username).Scan(&userID)
	if err != nil {
		return "User not found"
	}

	var reqExists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM requests WHERE user_id=? AND book_id=?)", userID, bookID).Scan(&reqExists)
	if err != nil {
		return "Internal Server Error 2"
	}

	if !reqExists {
		_, err = db.Exec("INSERT INTO requests (book_id, user_id, req_type, state) VALUES (?, ?, 'borrow', 'requested')", bookID, userID)
		if err != nil {
			return "Internal Server Error 3"
		}
		return ""
	} else {
		fmt.Println(err)
		return "Book Already existed"
	}
}
