package main

import (
	"database/sql"
	"final/cmd"
	"final/cmd/handlers"
	"final/cmd/repository"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "modernc.org/sqlite"
)

func main() {

	db := initDB("storage.db")
	// migrate(db)
	// createUsers(db, "admin1", "password1", func(password string) string {
	// 	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	// 	return string(bytes)
	// })
	// createUsers(db, "admin2", "password2", func(password string) string {
	// 	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	// 	return string(bytes)
	// })
	router := echo.New()
	repo := repository.NewRepository(db)
	apiH := handlers.API{StorageService: *repo}

	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		// This is a sample demonstration of how to attach middlewares in Echo
		return func(ctx echo.Context) error {
			log.Println("Echo middleware was called")
			return next(ctx)
		}
	})
	router.Use(middleware.BasicAuth(apiH.Authorize))

	// Add your handler (API endpoint) registrations here
	router.GET("/api", func(ctx echo.Context) error {
		return ctx.JSON(200, "Hello, World!")
	})
	router.GET("/api/list/export", apiH.GetCsv())
	router.GET("/api/weather", apiH.GetWeather())
	router.DELETE("/api/lists/:id", apiH.DeleteList())
	router.POST("/api/lists", apiH.PutList())
	router.GET("/api/lists", apiH.GetLists())
	router.DELETE("/api/tasks/:id", apiH.DeleteTask())
	router.POST("/api/lists/:id/tasks", apiH.PutTask())
	router.GET("/api/lists/:id/tasks", apiH.GetTasks())
	router.PATCH("/api/tasks/:id", apiH.PatchTasks())

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
func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS lists(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL
    ); 
	CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        text VARCHAR NOT NULL,
        listId INTEGER, 
        completed BOOLEAN,
		FOREIGN KEY(listId) REFERENCES lists(id) ON DELETE CASCADE
    ); 
	PRAGMA foreign_keys = ON;
    `

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
