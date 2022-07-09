package handlers

import (
	story "Lecture25/story"
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
	GetLastStoryTimeStamp() time.Time
	GetStories() []story.Story
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
		if api.StorageService.GetLastStoryTimeStamp().Before(treshold) {
			fmt.Println("Using API")
			stList = api.StoryService.GetStories(10)
		} else {
			fmt.Println("Using Storage")
			stList = api.StorageService.GetStories()
		}

		json.NewEncoder(w).Encode(stList)
	}
}
