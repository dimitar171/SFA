package main

import (
	handlers "Lecture25/handlers"
	repository "Lecture25/repository"
	"Lecture25/story"
	"database/sql"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {
	router := http.NewServeMux()
	db, err := sql.Open("sqlite", "dbHN.db")
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRepository(db)
	apiL := handlers.API{
		StorageService: repo,
		StoryService:   *story.NewStoryService("https://hacker-news.firebaseio.com", repo),
	}
	router.Handle("/api/top", apiL.HandleTopStories())
	log.Println("started")
	http.ListenAndServe(":8080", router)
}
