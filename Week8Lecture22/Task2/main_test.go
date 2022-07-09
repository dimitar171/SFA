package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestTopNames(t *testing.T) {
	//I made the test with only 1 stori, in order to be simple
	storiesNumber := 356925
	storieBody := top_stories{Title: "test", Score: 100}
	handlerString := fmt.Sprintf("/api/top/%d", storiesNumber)

	router := http.NewServeMux()
	router.Handle(handlerString, MockStories(storieBody))
	router.Handle("/", MockTopTen(storiesNumber))
	mockServer := httptest.NewServer(router)

	worker := NewWorker(mockServer.URL)
	fmt.Println(worker)
	got := worker.GetItem()
	want := top_stories{Title: "test", Score: 100}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`Got %v, want %v.`, got, want)
	}
}

//2 mocks, on for the stori number and one for the stori body( title and score )
func MockTopTen(numbers int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(numbers)
	}
}

func MockStories(s top_stories) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(s)
	}
}
