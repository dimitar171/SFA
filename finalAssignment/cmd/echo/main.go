package main

import (
	"database/sql"
	"final/cmd"
	"final/cmd/echo/rout"
	"final/cmd/handlers"
	"final/cmd/repository"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {

	db := initDB("storage.db")

	// createUsers(db, "admin1", "password1", func(password string) string {
	// 	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	// 	return string(bytes)
	// })

	repo := repository.NewRepository(db)
	apiH := handlers.API{StorageService: *repo}
	router := rout.Router(apiH)

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}

type HashDefinition func(string) string

func createUsers(db *sql.DB, user string, pass string, HashPassword HashDefinition) {
	sql := fmt.Sprintf("INSERT INTO users(username,password) VALUES('%s','%s')", user, HashPassword(pass))
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}
