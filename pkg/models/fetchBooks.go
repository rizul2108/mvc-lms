package models

import (
	_ "github.com/go-sql-driver/mysql"
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
		books = append(books, book)
	}

	return books, nil
}
