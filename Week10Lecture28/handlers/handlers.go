package handlers

import (
	"Lecture28/db"
	story "Lecture28/story"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type API struct {
	StorageService Storage
	StoryService   story.StoryService
}
type Storage interface {
	GetLastStoryTimeStamp(ctx context.Context) (time.Time, error)
	GetStories(ctx context.Context) ([]db.GetStoriesRow, error)
}

func (api API) HandleTopStories() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		treshold := time.Now().Add(-time.Hour)
		var stList []story.Story
		fmt.Println(treshold)
		tim, _ := api.StorageService.GetLastStoryTimeStamp(context.Background())
		if tim.Before(treshold) {
			fmt.Println("Using API")
			stList = api.StoryService.GetStories(10)
		} else {
			fmt.Println("Using Storage")
			stList, _ := api.StorageService.GetStories(context.Background())
			json.NewEncoder(w).Encode(stList)
		}
		json.NewEncoder(w).Encode(stList)
	}
}
