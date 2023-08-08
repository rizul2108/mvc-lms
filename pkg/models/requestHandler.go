package models

import (
	"fmt"
	"time"
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
		currentTime := time.Now()
		_, err = db.Exec("INSERT INTO requests (book_id, user_id, request_type, state,request_date) VALUES (?, ?, 'Borrow', 'Requested',?)", bookID, userID, currentTime)
		if err != nil {
			return "Internal Server Error 3"
		}
		return ""
	} else {
		fmt.Println(err)
		return "Book Already existed"
	}
}

func DeleteRequest(requestId int) string {
	db, err := Connection()
	if err != nil {
		fmt.Println(err)
		return "Error in conncting to database"
	}
	var reqType string

	error := db.QueryRow("SELECT request_type FROM requests WHERE request_id = ?", requestId).Scan(&reqType)
	if error != nil {
		return "Internal Server Error"
	}
	if reqType == "Borrow" {
		result, err := db.Exec("DELETE FROM requests WHERE request_id = ?", requestId)
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
		result, err := db.Exec(`UPDATE requests SET request_type="Accepted", state="Owned" WHERE request_id =?`, requestId)
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

func ReturnBook(requestId int) string {
	db, err := Connection()
	if err != nil {
		fmt.Println(err)
		return "Error in connecting to db"
	}
	result, err := db.Exec(`UPDATE requests SET state="Requested" , request_type="return" WHERE request_id =?`, requestId)
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

func AcceptRequest(requestId int) string {
	db, err := Connection()
	if err != nil {
		fmt.Println(err)
		return "Error in connecting to db"
	}
	var reqType string
	var bookID int
	error := db.QueryRow("SELECT request_type,book_id FROM requests WHERE request_id = ?", requestId).Scan(&reqType, &bookID)
	if error != nil {
		return "Internal Server Error"
	}
	if reqType == "Borrow" {
		currentTime := time.Now()
		result, err := db.Exec(`UPDATE requests SET request_type="Accepted", state="Owned",request_date=? WHERE request_id = ?`, currentTime, requestId)
		if err != nil {
			fmt.Println(err)
			return "Internal Server error"
		} else {
			rowsAffected, err := result.RowsAffected()
			if rowsAffected > 0 {
				db.Exec(`UPDATE books SET quantity=quantity - 1 WHERE book_id=?`, bookID)
				return ""
			} else if err != nil || rowsAffected == 0 {
				fmt.Println(err)
				return "request doesn't exist"
			}
		}
	} else {
		result, err := db.Exec(`DELETE FROM requests WHERE request_id=?`, requestId)
		if err != nil {
			fmt.Println(err)
			return "Internal Server error"
		} else {
			rowsAffected, err := result.RowsAffected()
			if rowsAffected > 0 {
				db.Exec(`UPDATE books SET quantity=quantity + 1 WHERE book_id=?`, bookID)
				return ""
			} else if err != nil || rowsAffected == 0 {
				fmt.Println(err)
				return "request doesn't exist"
			}
		}
	}
	return "error"
}

func DeclineRequest(requestId int) string {
	db, err := Connection()
	if err != nil {
		fmt.Println(err)
		return "Error in connecting to db"
	}
	result, err := db.Exec("DELETE FROM requests WHERE request_id = ?", requestId)
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

func DeclineAdminReq(userId int) string {
	db, err := Connection()
	if err != nil {
		fmt.Println(err)
		return "Error in connecting to db"
	}
	result, err := db.Exec(`UPDATE users SET type="client" WHERE user_id = ?`, userId)
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
	return ""
}

func AcceptAdminReq(userId int) string {
	db, err := Connection()
	if err != nil {
		fmt.Println(err)
		return "Error in connecting to db"
	}
	result, err := db.Exec(`UPDATE users SET type="admin" WHERE user_id = ?`, userId)
	if err != nil {
		fmt.Println(err)
		return "Internal Server error"
	} else {
		rowsAffected, err := result.RowsAffected()
		if rowsAffected > 0 {
			db.Exec(`DELETE FROM requests WHERE user_id = ?`, userId)
			return ""
		} else if err != nil || rowsAffected == 0 {
			fmt.Println(err)
			return "request doesn't exist"
		}
	}
	return ""
}
