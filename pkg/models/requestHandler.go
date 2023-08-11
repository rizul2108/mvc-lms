package models

import (
	"time"
)

func AddRequest(bookID int, username string) string {
	db, err := Connection()
	if err != nil {
		return "Internal Server Error 1"
	}

	var userID int
	err = db.QueryRow("SELECT userID FROM users WHERE username=?", username).Scan(&userID)
	if err != nil {
		return "User not found"
	}

	var reqExists bool
	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM requests WHERE userID=? AND bookID=?)", userID, bookID).Scan(&reqExists)
	if err != nil {
		return "Internal Server Error 2"
	}

	if !reqExists {
		currentTime := time.Now()
		_, err = db.Exec("INSERT INTO requests (bookID, userID, requestType, state,requestDate) VALUES (?, ?, 'Borrow', 'Requested',?)", bookID, userID, currentTime)
		if err != nil {
			return "Internal Server Error 3"
		}
		return ""
	} else {
		return "Book Already existed"
	}
}

func DeleteRequest(requestId int) string {
	db, err := Connection()
	if err != nil {
		return "Error in conncting to database"
	}
	var reqType string

	error := db.QueryRow("SELECT requestType FROM requests WHERE requestID = ?", requestId).Scan(&reqType)
	if error != nil {
		return "Internal Server Error"
	}
	if reqType == "Borrow" {
		result, err := db.Exec("DELETE FROM requests WHERE requestID = ?", requestId)
		if err != nil {
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
		result, err := db.Exec(`UPDATE requests SET requestType="Accepted", state="Owned" WHERE requestID =?`, requestId)
		if err != nil {
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
		return "Error in connecting to db"
	}
	result, err := db.Exec(`UPDATE requests SET state="Requested" , requestType="return" WHERE requestID =?`, requestId)
	if err != nil {

		return "Internal Server error"
	} else {
		rowsAffected, err := result.RowsAffected()
		if rowsAffected > 0 {
			return ""
		} else if err != nil || rowsAffected == 0 {

			return "request doesn't exist"
		}
	}
	return "error"
}

func AcceptRequest(requestId int) string {
	db, err := Connection()
	if err != nil {

		return "Error in connecting to db"
	}
	var reqType string
	var bookID int
	error := db.QueryRow("SELECT requestType,bookID FROM requests WHERE requestID = ?", requestId).Scan(&reqType, &bookID)
	if error != nil {
		return "Internal Server Error"
	}
	if reqType == "Borrow" {
		currentTime := time.Now()
		result, err := db.Exec(`UPDATE requests SET requestType="Accepted", state="Owned",requestDate=? WHERE requestID = ?`, currentTime, requestId)
		if err != nil {

			return "Internal Server error"
		} else {
			rowsAffected, err := result.RowsAffected()
			if rowsAffected > 0 {
				db.Exec(`UPDATE books SET quantity=quantity - 1 WHERE bookID=?`, bookID)
				return ""
			} else if err != nil || rowsAffected == 0 {

				return "request doesn't exist"
			}
		}
	} else {
		result, err := db.Exec(`DELETE FROM requests WHERE requestID=?`, requestId)
		if err != nil {

			return "Internal Server error"
		} else {
			rowsAffected, err := result.RowsAffected()
			if rowsAffected > 0 {
				db.Exec(`UPDATE books SET quantity=quantity + 1 WHERE bookID=?`, bookID)
				return ""
			} else if err != nil || rowsAffected == 0 {

				return "request doesn't exist"
			}
		}
	}
	return "error"
}

func DeclineRequest(requestId int) string {
	db, err := Connection()
	if err != nil {

		return "Error in connecting to db"
	}
	result, err := db.Exec("DELETE FROM requests WHERE requestID = ?", requestId)
	if err != nil {

		return "Internal Server error"
	} else {
		rowsAffected, err := result.RowsAffected()
		if rowsAffected > 0 {
			return ""
		} else if err != nil || rowsAffected == 0 {

			return "request doesn't exist"
		}
	}
	return "error"
}

func DeclineAdminReq(userId int) string {
	db, err := Connection()
	if err != nil {

		return "Error in connecting to db"
	}
	result, err := db.Exec(`UPDATE users SET type="client" WHERE userID = ?`, userId)
	if err != nil {

		return "Internal Server error"
	} else {
		rowsAffected, err := result.RowsAffected()
		if rowsAffected > 0 {
			return ""
		} else if err != nil || rowsAffected == 0 {

			return "request doesn't exist"
		}
	}
	return ""
}

func AcceptAdminReq(userId int) string {
	db, err := Connection()
	if err != nil {

		return "Error in connecting to db"
	}
	result, err := db.Exec(`UPDATE users SET type="admin" WHERE userID = ?`, userId)
	if err != nil {

		return "Internal Server error"
	} else {
		rowsAffected, err := result.RowsAffected()
		if rowsAffected > 0 {
			db.Exec(`DELETE FROM requests WHERE userID = ?`, userId)
			return ""
		} else if err != nil || rowsAffected == 0 {

			return "request doesn't exist"
		}
	}
	return ""
}
