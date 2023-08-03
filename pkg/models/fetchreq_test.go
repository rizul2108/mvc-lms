package models

import (
	_ "github.com/go-sql-driver/mysql"
	"mvc-go/pkg/types"
	"reflect"
	"testing"
)

func TestFetchRequests(t *testing.T) {
	gotRequests, gotStr := FetchRequests("admin")
	var requests []types.Request
	str := ""
	wantRequests := requests
	wantStr := str

	if !reflect.DeepEqual(gotRequests, wantRequests) || gotStr != wantStr {
		t.Errorf("got (%v, %q), wanted (%v, %q)", gotRequests, gotStr, wantRequests, wantStr)
	}
}
