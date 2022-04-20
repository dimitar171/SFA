package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type top_stories struct {
	Title string `json:"title"`
	Score int    `json:"score"`
}

type Worker struct {
	serverURL string
}

func NewWorker(serverUrl string) *Worker {
	return &Worker{serverUrl}
}

//we get the top news number
func (w Worker) GetTop() int {
	var arr int
	respTop, _ := http.Get(w.serverURL)
	json.NewDecoder(respTop.Body).Decode(&arr)
	respTop.Body.Close()
	fmt.Println(arr)
	return arr
}

//we get storie title and score depending on the number
func (w Worker) GetItem() top_stories {
	arr := w.GetTop()
	topStori := fmt.Sprintf(w.serverURL+"/api/top/%d", arr)
	resp, _ := http.Get(topStori)
	payload := top_stories{}
	json.NewDecoder(resp.Body).Decode(&payload)
	resp.Body.Close()
	return payload
}
