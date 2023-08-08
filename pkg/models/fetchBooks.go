package models

import (
	"mvc-go/pkg/types"
)

func FetchBooks() ([]types.Book, error) {
	db, err := Connection()
	rows, err := db.Query("SELECT book_id, title, author, quantity FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Quantity)
		if err != nil {
			return nil, err
		}
		db.QueryRow(`SELECT COUNT(*) FROM requests WHERE book_id =?`, book.ID).Scan(&book.IssuedQuantity)

		books = append(books, book)
	}

	return books, nil
}
