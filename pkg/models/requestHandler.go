package models

import (
	"fmt"
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

func DeleteRequest(reqID int) string {
	db, err := Connection()
	if err != nil {
		fmt.Println(err)
		return "Error in conncting to database"
	}
	var reqType string

	error := db.QueryRow("SELECT req_type FROM requests WHERE request_id = ?", reqID).Scan(&reqType)
	if error != nil {
		return "Internal Server Error"
	}
	if reqType == "borrow" {
		result, err := db.Exec("DELETE FROM requests WHERE request_id = ?", reqID)
		if err != nil {
			fmt.Println(err)
			return "Internal Server error"
		} else {
			rowsAffected, err := result.RowsAffected()
			if rowsAffected > 0 {
				return ""
			} else if err != nil || rowsAffected == 0 {
				return "request doesn't exist"
			}
		}
	} else {
		result, err := db.Exec(`UPDATE requests SET req_type="accepted", state="owned" WHERE request_id =?`, reqID)
		if err != nil {
			fmt.Println(err)
			return "Internal Server error"
		} else {
			rowsAffected, err := result.RowsAffected()
			if rowsAffected > 0 {
				return ""
			} else if err != nil || rowsAffected == 0 {
				return "request doesn't exist"
			}
		}
	}
	return "Error"
}

func ReturnBook(reqID int) string {
	db, err := Connection()
	if err != nil {
		fmt.Println(err)
		return "Error in connecting to db"
	}
	result, err := db.Exec(`update requests set state="requested" , req_type="return" WHERE request_id =?`, reqID)
	if err != nil {
		fmt.Println(err)
		return "Internal Server error"
	} else {
		rowsAffected, err := result.RowsAffected()
		if rowsAffected > 0 {
			return ""
		} else if err != nil || rowsAffected == 0 {
			fmt.Println(err)
			return "request doesn't exist"
		}
	}
	return "error"
}

func AcceptRequest(reqID int) string {
	db, err := Connection()
	if err != nil {
		fmt.Println(err)
		return "Error in connecting to db"
	}
	var reqType string
	var bookID int
	error := db.QueryRow("SELECT req_type,book_id FROM requests WHERE request_id = ?", reqID).Scan(&reqType, &bookID)
	if error != nil {
		return "Internal Server Error"
	}
	if reqType == "borrow" {
		result, err := db.Exec(`update requests set req_type="accepted", state="owned" WHERE request_id = ?`, reqID)
		if err != nil {
			fmt.Println(err)
			return "Internal Server error"
		} else {
			rowsAffected, err := result.RowsAffected()
			if rowsAffected > 0 {
				db.Exec(`update books set quantity=quantity - 1 where book_id=?`, bookID)
				return ""
			} else if err != nil || rowsAffected == 0 {
				fmt.Println(err)
				return "request doesn't exist"
			}
		}
	} else {
		result, err := db.Exec(`admin/deleteBooks from requests where request_id=?`, reqID)
		if err != nil {
			fmt.Println(err)
			return "Internal Server error"
		} else {
			rowsAffected, err := result.RowsAffected()
			if rowsAffected > 0 {
				db.Exec(`update books set quantity=quantity + 1 where book_id=?`, bookID)
				return ""
			} else if err != nil || rowsAffected == 0 {
				fmt.Println(err)
				return "request doesn't exist"
			}
		}
	}
	return "error"
}
func DeclineRequest(reqID int) string {
	db, err := Connection()
	if err != nil {
		fmt.Println(err)
		return "Error in connecting to db"
	}
	result, err := db.Exec("DELETE FROM requests WHERE request_id = ?", reqID)
	if err != nil {
		fmt.Println(err)
		return "Internal Server error"
	} else {
		rowsAffected, err := result.RowsAffected()
		if rowsAffected > 0 {
			return ""
		} else if err != nil || rowsAffected == 0 {
			fmt.Println(err)
			return "request doesn't exist"
		}
	}
	return "error"
}
