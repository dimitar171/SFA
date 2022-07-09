package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

type top_stories struct {
	Title string `json:"title"`
	Score int    `json:"score"`
}
type myArray struct {
	PageTitle   string
	Top_stories []top_stories
}
type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func HandleUserTop(stories []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(stories)
	}

}
func HandleHttp(stories myArray, tmp *template.Template) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		data := stories
		tmp.Execute(w, data)
	}
}

func generator(arr []int) chan top_stories {
	c := make(chan top_stories)
	go func() {
		for _, result := range arr {
			topStori := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", result)
			resp, _ := http.Get(topStori)

			payload := top_stories{}
			json.NewDecoder(resp.Body).Decode(&payload)
			resp.Body.Close()
			c <- payload
		}
		close(c)
	}()
	return c
}

func main() {
	var arr []int

	//Get top stories number

	respTop, _ := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")

	json.NewDecoder(respTop.Body).Decode(&arr)
	arr = arr[0:10]
	respTop.Body.Close()
	var rez []top_stories
	//Get title and score with generator function concurently
	for v := range generator(arr) {
		rez = append(rez, v)
	}
	//create JSON response
	res := myArray{"Hacker news top 10", rez}
	result, _ := json.MarshalIndent(res, "", "")

	//show the results on localhost:9000/top
	tmp, _ := template.ParseFiles("my_template.html")
	mux := http.NewServeMux()
	mux.Handle("/api/top", HandleUserTop(result))
	mux.Handle("/top", HandleHttp(res, tmp))
	http.ListenAndServe(":9000", mux)
}
