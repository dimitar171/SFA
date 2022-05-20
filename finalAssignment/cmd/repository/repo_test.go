package repository

import (
	"database/sql"
	"reflect"
	"testing"

	_ "modernc.org/sqlite"
)

const (
	createListTable    = " CREATE TABLE IF NOT EXISTS lists(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,name VARCHAR NOT NULL,userId INTEGER); "
	createTaskTable    = "CREATE TABLE IF NOT EXISTS tasks(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,text VARCHAR NOT NULL, listId INTEGER, completed BOOLEAN)"
	insertList         = "INSERT INTO lists(name,userId) VALUES(?,?)"
	insertTask         = "INSERT INTO Tasks(text,listId,completed) VALUES(?,?,?)"
	insertUser         = "INSERT INTO users(username,password) VALUES(?,?)"
	createUserTable    = "CREATE TABLE IF NOT EXISTS users(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,username VARCHAR NOT NULL,password VARCHAR NOT NULL); "
	createCurrentTable = "CREATE TABLE IF NOT EXISTS currentUsser(id INTEGER NOT NULL PRIMARY KEY,Usser INTEGER NOT NULL);INSERT INTO currentUsser(Usser) VALUES(1)"
)

func TestGetLists(t *testing.T) {

	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in memory DB")
	}
	_, err = mockDb.Exec(createListTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}
	_, err = mockDb.Exec(createCurrentTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}
	repo := NewRepository(mockDb)

	_, err = mockDb.Exec(insertList, "UnitTest", 1)
	if err != nil {
		t.Fatal(err)
	}
	res := repo.GetLists()
	wantedLists := []List{{ID: 1, Name: "UnitTest", UserId: 1}}
	if !reflect.DeepEqual(res, wantedLists) {
		t.Fatal("Failed to get lists")
	}
}

func TestPutLists(t *testing.T) {
	mockList := List{ID: 1, Name: "UnitTest", UserId: 1}

	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in memory DB")
	}
	_, err = mockDb.Exec(createListTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}
	_, err = mockDb.Exec(createCurrentTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}
	repo := NewRepository(mockDb)
	repo.PutList("UnitTest")
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
	_, err = mockDb.Exec(createCurrentTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}
	repo := NewRepository(mockDb)

	_, err = mockDb.Exec(insertList, "UnitTest", 1)
	if err != nil {
		t.Fatal(err)
	}
	repo.DeleteList(1)
	res := repo.GetLists()
	wantedLists := []List{}
	if !reflect.DeepEqual(res, wantedLists) {
		t.Fatal("Failed to delete List")
	}
}
func TestGetTask(t *testing.T) {

	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in memory DB")
	}
	_, err = mockDb.Exec(createTaskTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}

	repo := NewRepository(mockDb)

	_, err = mockDb.Exec(insertTask, "UnitTest", 1, false)
	if err != nil {
		t.Fatal(err)
	}
	res := repo.GetTasks(1)
	wantedLists := []Task{{Id: 1, Text: "UnitTest", ListId: 1, Completed: false}}
	if !reflect.DeepEqual(res, wantedLists) {
		t.Fatal("Failed to get lists")
	}
}

func TestPutTask(t *testing.T) {
	mockList := Task{Id: 1, Text: "UnitTest", ListId: 1, Completed: false}

	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in memory DB")
	}
	_, err = mockDb.Exec(createTaskTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}

	repo := NewRepository(mockDb)
	repo.PutTask("UnitTest", 1, false)
	res := repo.GetTasks(1)
	if !reflect.DeepEqual(res[0], mockList) {
		t.Fatal("Failed to put List")
	}
}

func TestDeleteTask(t *testing.T) {
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in memory DB")
	}
	_, err = mockDb.Exec(createTaskTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}

	repo := NewRepository(mockDb)

	_, err = mockDb.Exec(insertTask, "UnitTest", 1, false)
	if err != nil {
		t.Fatal(err)
	}

	repo.DeleteTask(1)
	res := repo.GetTasks(1)
	wantedTask := []Task{}
	if !reflect.DeepEqual(res, wantedTask) {
		t.Fatal("Failed to delete List")
	}
}

func TestPatchTask(t *testing.T) {
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in memory DB")
	}
	_, err = mockDb.Exec(createTaskTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}
	repo := NewRepository(mockDb)
	_, err = mockDb.Exec(insertTask, "UnitTest", 1, false)
	if err != nil {
		t.Fatal(err)
	}
	repo.PatchTask(1)
	res := repo.GetTasks(1)
	wantedLists := []Task{{Id: 1, Text: "UnitTest", ListId: 1, Completed: true}}
	if !reflect.DeepEqual(res, wantedLists) {
		t.Fatal("Failed to delete List")
	}
}

func TestPatchCurrent(t *testing.T) {
	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in memory DB")
	}
	_, err = mockDb.Exec(createCurrentTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}
	repo := NewRepository(mockDb)

	repo.PatchCurrentUser(3)
	res := repo.GetCurrentUser()
	wantedLists := 3
	if !reflect.DeepEqual(res, wantedLists) {
		t.Fatal("Failed to delete List")
	}
}

func TestGetUsers(t *testing.T) {

	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in memory DB")
	}
	_, err = mockDb.Exec(createUserTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}

	repo := NewRepository(mockDb)
	_, err = mockDb.Exec(insertUser, "UnitTest", "UnitTest")
	if err != nil {
		t.Fatal(err)
	}
	res := repo.GetUsers()
	wantedUser := []User{{ID: 1, Name: "UnitTest", Pass: "UnitTest"}}
	if !reflect.DeepEqual(res, wantedUser) {
		t.Fatal("Failed to get lists")
	}
}

func TestGetCurrent(t *testing.T) {

	mockDb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("Failed to open in memory DB")
	}
	_, err = mockDb.Exec(createCurrentTable)
	if err != nil {
		t.Fatal("Failed to create table")
	}

	repo := NewRepository(mockDb)

	res := repo.GetCurrentUser()
	wantedCurrent := 1
	if !reflect.DeepEqual(res, wantedCurrent) {
		t.Fatal("Failed to get lists")
	}
}
