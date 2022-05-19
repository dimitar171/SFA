package repository

import (
	"database/sql"
	"final/cmd/todos"
	"fmt"
	"reflect"
	"testing"

	_ "modernc.org/sqlite"
)

const (
	createListTable = " CREATE TABLE IF NOT EXISTS lists(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,name VARCHAR NOT NULL); "
	insertList      = "INSERT INTO lists(name) VALUES(?)"
)

func TestGetLists(t *testing.T) {

	t.Errorf("The code did not panic")
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in memory DB")
	}
	_, err = mockDb.Exec(createListTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}
	repo := NewRepository(mockDb)

	_, err = mockDb.Exec(insertList, "UnitTest")
	if err != nil {
		t.Fatal(err)
	}
	res := repo.GetLists()
	wantedLists := []todos.List{{ID: 1, Name: "UnitTest"}}
	if !reflect.DeepEqual(res, wantedLists) {
		t.Fatal("Failed to get lists")
	}
}

func TestPutLists(t *testing.T) {
	mockList := todos.List{ID: 1, Name: "UnitTest"}

	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in memory DB")
	}
	_, err = mockDb.Exec(createListTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}
	repo := NewRepository(mockDb)
	repo.PutList(mockList.Name)
	res := repo.GetLists()
	if !reflect.DeepEqual(res[0], mockList) {
		t.Fatal("Failed to put List")
	}
}

func TestDeleteLists(t *testing.T) {
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in memory DB")
	}
	_, err = mockDb.Exec(createListTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}
	repo := NewRepository(mockDb)

	_, err = mockDb.Exec(insertList, "UnitTest")
	if err != nil {
		t.Fatal(err)
	}
	repo.DeleteList(1)
	res := []todos.List{{ID: 0}}
	wantedLists := []todos.List{{}}
	fmt.Println(res)
	fmt.Println(wantedLists)
	if !reflect.DeepEqual(res, wantedLists) {
		t.Fatal("Failed to delete List")
	}
}
