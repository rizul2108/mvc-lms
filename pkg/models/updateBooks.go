package models

import (
	"strconv"
)

func AddQuantity(id, quantity string) string {
	bookID, err := strconv.Atoi(id)
	qty, err := strconv.Atoi(quantity)
	db, err := Connection()
	_, err = db.Exec("UPDATE books SET quantity = quantity + ? WHERE bookID = ?", qty, bookID)
	if err != nil {
		return "Error executing the UPDATE query"
	}
	return ""
}

func DecreaseQuantity(id, quantity string) string {
	bookID, err := strconv.Atoi(id)
	qty, err := strconv.Atoi(quantity)
	db, err := Connection()
	_, err = db.Exec("UPDATE books SET quantity = quantity - ? WHERE bookID = ?", qty, bookID)
	if err != nil {
		return "Error executing the UPDATE query"
	}
	return ""
}

func DeleteBook(id string) string {
	bookID, err := strconv.Atoi(id)
	db, err := Connection()
	_, err = db.Exec("DELETE FROM books WHERE bookID = ?", bookID)
	if err != nil {
		return "Error executing the UPDATE query"
	}
	return ""
}
