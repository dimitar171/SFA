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
type myArray struct {
	Top_stories []top_stories
}

func HandleUserTop(stories []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(stories)
	}

}
func main() {
	var arr []int

	//make a call for top stories

	respTop, _ := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty")

	json.NewDecoder(respTop.Body).Decode(&arr)
	arr = arr[0:10]
	respTop.Body.Close()
	// get the first 10
	var rez []top_stories

	// c := make(chan int)
	// go func() {
	// for i:= 0; i < 10 ; i++ {
	// 	c <- i
	//  }
	// 	close(c)
	// }()
	// return c
	for _, result := range arr {
		topStori := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty", result)
		resp, _ := http.Get(topStori)

		payload := top_stories{}
		json.NewDecoder(resp.Body).Decode(&payload)

		rez = append(rez, payload)
		resp.Body.Close()
	}
	res := myArray{rez}
	result, _ := json.MarshalIndent(res, "", "")

	//create a server
	mux := http.NewServeMux()
	mux.Handle("/top", HandleUserTop(result))
	http.ListenAndServe(":9000", mux)
}
