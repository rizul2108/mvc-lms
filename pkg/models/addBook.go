package models

import (
	"log"
)

func AddBook(title, author string, quantity int) string {

	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM books WHERE title=? AND author=?", title, author)
	if err != nil {
		return "There was error"
	}
	defer rows.Close()

	if rows.Next() {
		_, err = db.Exec("UPDATE books SET quantity=quantity+? WHERE title=?", quantity, title)
		if err != nil {
			return "Database error"
		}
	} else {
		_, err = db.Exec("INSERT INTO books (title, author, quantity) VALUES (?, ?, ?)", title, author, quantity)
		if err != nil {
			return "Database error"
		}
	}
	return ""
}
