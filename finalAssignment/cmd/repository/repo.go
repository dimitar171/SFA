package repository

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type List struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserId int    `json:"userId"`
}
type Task struct {
	Id        int    `json:"id"`
	Text      string `json:"text"`
	ListId    int    `json:"listId"`
	Completed bool   `json:"completed"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (rp *Repository) GetLists() []List {
	id := rp.GetCurrentUser()
	sql := fmt.Sprintf("SELECT * FROM lists WHERE userId=%d", id)
	rows, err := rp.db.Query(sql)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	result := []List{}
	for rows.Next() {
		temp := List{}
		err2 := rows.Scan(&temp.ID, &temp.Name, &temp.UserId)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, temp)
	}
	return result
}

func (rp *Repository) PutList(name string) (int64, error) {
	id := rp.GetCurrentUser()
	sql := "INSERT INTO lists(name,userId) VALUES(?,?)"
	stmt, err := rp.db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(name, id)
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func (rp *Repository) DeleteList(id int) (int64, error) {
	sql := "DELETE FROM Lists WHERE id = ?"
	stmt, err := rp.db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	result, err2 := stmt.Exec(id)
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}

func (rp *Repository) GetTasks(id int) []Task {
	sql := fmt.Sprintf("SELECT * FROM tasks WHERE listId=%d", id)
	rows, err := rp.db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := []Task{}
	for rows.Next() {
		temp := Task{}
		err2 := rows.Scan(&temp.Id, &temp.Text, &temp.ListId, &temp.Completed)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, temp)
	}
	return result
}

func (rp *Repository) PutTask(text string, listId int, completed bool) (int64, error) {
	sql := "INSERT INTO Tasks(text,listId,completed) VALUES(?,?,?)"
	stmt, err := rp.db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(text, listId, completed)
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func (rp *Repository) DeleteTask(id int) (int64, error) {
	sql := "DELETE FROM Tasks WHERE id = ?"
	stmt, err := rp.db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	result, err2 := stmt.Exec(id)
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}

func (rp *Repository) PatchTask(id int) (int64, error) {
	sql := "UPDATE Tasks SET completed='TRUE' WHERE id=?"
	stmt, err := rp.db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	result, err2 := stmt.Exec(id)
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}

func (rp *Repository) GetUsers() []User {
	sql := "SELECT * FROM users"
	rows, err := rp.db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	result := []User{}
	for rows.Next() {
		temp := User{}
		err2 := rows.Scan(&temp.ID, &temp.Name, &temp.Pass)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, temp)
	}
	return result
}

func (rp *Repository) PatchCurrentUser(id int) {
	sql := fmt.Sprintf("UPDATE currentUsser SET Usser=%d WHERE id=1", id)
	stmt, err := rp.db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	_, err2 := stmt.Exec()
	if err2 != nil {
		panic(err2)
	}
}

func (rp *Repository) GetCurrentUser() int {
	sql := "SELECT Usser FROM currentUsser"
	rows, err := rp.db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	var result int
	for rows.Next() {
		err2 := rows.Scan(&result)
		if err2 != nil {
			panic(err2)
		}
	}
	return result
}
