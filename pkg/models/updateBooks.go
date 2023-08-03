package models

import (
	"strconv"
)

func AddQuantity(id, quantity string) string {
	bookID, err := strconv.Atoi(id)
	qty, err := strconv.Atoi(quantity)
	db, err := Connection()
	_, err = db.Exec("UPDATE books SET quantity = quantity + ? WHERE book_id = ?", qty, bookID)
	if err != nil {
		return "Error executing the update query"
	}
	return ""
}
func DecreaseQuantity(id, quantity string) string {
	bookID, err := strconv.Atoi(id)
	qty, err := strconv.Atoi(quantity)
	db, err := Connection()
	_, err = db.Exec("UPDATE books SET quantity = quantity - ? WHERE book_id = ?", qty, bookID)
	if err != nil {
		return "Error executing the update query"
	}
	return ""
}
