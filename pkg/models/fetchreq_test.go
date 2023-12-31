package models

import (
	"mvc-go/pkg/types"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestFetchRequests(t *testing.T) {
	// Create a mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	userID := 1
	username := "testuser"
	rows := sqlmock.NewRows([]string{"userID"}).AddRow(userID)
	mock.ExpectQuery("SELECT userID FROM users WHERE username=?").
		WithArgs(username).
		WillReturnRows(rows)

	currentTime := time.Now()
	dateString := currentTime.Format("2006-01-02 15:04:05")

	requestRows := sqlmock.NewRows([]string{"requestID", "bookID", "state", "requestType", "requestDate"}).
		AddRow(1, 101, "Requested", "Borrow", dateString)
	mock.ExpectQuery("SELECT requestID, bookID, state, requestType, requestDate FROM requests WHERE userID=?").
		WithArgs(userID).
		WillReturnRows(requestRows)

	bookID := 101
	bookTitleRow := sqlmock.NewRows([]string{"title"}).AddRow("Book Title 1")
	mock.ExpectQuery("SELECT title FROM books WHERE bookID=?").
		WithArgs(bookID).
		WillReturnRows(bookTitleRow)
	currentTime, err = time.Parse("2006-01-02 15:04:05", dateString)
	requests, errStr := FetchRequests(db, username)

	// Check the results
	wantRequests := []types.Request{
		{RequestID: 1, BookID: 101, State: "Requested", RequestType: "Borrow", BookTitle: "Book Title 1", RequestDate: currentTime, RequestDateString: dateString, Fine: 0},
	}
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
