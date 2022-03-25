package main

import (
	"flag"
	"log"
	"net/http"
)

//try this
//.\multiping.exe -c 5 https://google.com https://facebook.com https://sitedoesnotexist.com https://www.instagram.com/ https://www.udemy.com/
func main() {
	var urls []string
	var concurency int

	flag.IntVar(&concurency, "c", 2, "number of concurrencies")
	flag.Parse()
	urls = flag.Args()

	fetchURL(urls, concurency)

}
func fetchURL(urls []string, concurency int) {
	processQueue := make(chan string, concurency)
	done := make(chan string)
	go func() {
		for _, urlToProcess := range urls {
			processQueue <- urlToProcess
			go func(url string) {
				err := pingURL(url)
				if err != nil {
					done <- "Error reaching the site " + url
				} else {
					done <- url + " ok"
				}
				<-processQueue
			}(urlToProcess)
		}
	}()
	for range urls {
		log.Println("Done: ", <-done)
	}
}

func pingURL(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("Got response for %s: %d\n", url, resp.StatusCode)
	return nil
}
