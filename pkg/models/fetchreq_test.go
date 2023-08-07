package models

import (
	"github.com/DATA-DOG/go-sqlmock"
	"mvc-go/pkg/types"
	"reflect"
	"testing"
)

func TestFetchRequests(t *testing.T) {
	// Create a mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Define the expected query and results for username query
	userID := 1
	username := "testuser"
	rows := sqlmock.NewRows([]string{"user_id"}).AddRow(userID)
	mock.ExpectQuery("SELECT user_id FROM users WHERE username=?").
		WithArgs(username).
		WillReturnRows(rows)

	// Define the expected query and results for requests
	requestRows := sqlmock.NewRows([]string{"request_id", "book_id", "state", "req_type"}).
		AddRow(1, 101, "Pending", "Borrow")
	mock.ExpectQuery("SELECT request_id,book_id, state, req_type FROM requests where user_id=?").
		WithArgs(userID).
		WillReturnRows(requestRows)

	// Define the expected query and result for book titles
	bookID := 101
	bookTitleRow := sqlmock.NewRows([]string{"title"}).AddRow("Book Title 1")
	mock.ExpectQuery("SELECT title from books where book_id=?").
		WithArgs(bookID).
		WillReturnRows(bookTitleRow)

	// Call the function to be tested directly with the mock db object
	requests, errStr := FetchRequests(db, username)

	// Check the results
	wantRequests := []types.Request{
		{RequestID: 1, BookID: 101, State: "Pending", RequestType: "Borrow", BookTitle: "Book Title 1"}}
	if !reflect.DeepEqual(requests, wantRequests) {
		t.Errorf("got %v, wanted %v", requests, wantRequests)
	}

	if errStr != "" {
		t.Errorf("unexpected error: %s", errStr)
	}

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
