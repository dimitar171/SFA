package story

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func handleTopStories(ids []int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		json.NewEncoder(w).Encode(ids)
	}
}
func handleGetStory(stories []Story) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("bla")
		id := 10
		var resultStory Story
		for _, v := range stories {
			if v.Id == id {
				fmt.Println("blabla")
				resultStory = v
				break
			}
		}
		json.NewEncoder(w).Encode(resultStory)
	}
}

type FakeStorage struct {
	savedStories []Story
}

func (fs *FakeStorage) SaveStories(sList []Story) {
	fmt.Println("IN FAKE SaveStories")
	fs.savedStories = sList
}

func TestTopStoriesIds(t *testing.T) {
	router := http.NewServeMux()
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	router.Handle("/v0/topstories.json", handleTopStories(ids))
	mockServer := httptest.NewServer(router)

	wont := ids[:10]
	ss := NewStoryService(mockServer.URL, &FakeStorage{})
	got := ss.GetTopStoriesIds(10)
	if !reflect.DeepEqual(got, wont) {
		t.Fatalf("got: %v, want %v", got, wont)
	}
}

func TestTopStories(t *testing.T) {
	router := http.NewServeMux()
	ids := []int{10}
	stories := []Story{
		{
			Id:    10,
			Title: "test 10",
			Score: 15,
		},
	}
	router.Handle("/v0/topstories.json", handleTopStories(ids))
	urlTest := fmt.Sprintf("/v0/item/%d.json", ids[0])
	router.Handle(urlTest, handleGetStory(stories))
	mockServer := httptest.NewServer(router)
	wont := stories
	fs := &FakeStorage{}
	ss := NewStoryService(mockServer.URL, fs)
	got := ss.GetStories(1)

	if !reflect.DeepEqual(got, wont) {
		t.Fatalf("got: %v, want %v", got, wont)
	}
	if !reflect.DeepEqual(got, fs.savedStories) {
		t.Fatalf("got: %v, want %v", fs.savedStories, got)
	}
}
