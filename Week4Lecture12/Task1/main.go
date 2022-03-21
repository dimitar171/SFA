package main

import (
	"log"

	// "sync"
	"time"
)

func main() {

	out := generateThrottled("foo", 2, time.Second)

	for f := range out {
		log.Println(f)
	}

}

func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {
	evenQueue := make(chan string, bufferLimit)
	outChan := make(chan string)

	go func() {
		for {
			timeoutChan := time.After(clearInterval)
			select {
			case evenQueue <- data:
				outChan <- data
			case <-timeoutChan:
				for i := 0; i < bufferLimit; i++ {

					<-evenQueue
				}
			}
		}
	}()

	return outChan
}
