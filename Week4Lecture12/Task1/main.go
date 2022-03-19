package main

import (
	"log"
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

	go func() {
		for {
			timeoutChan := time.After(clearInterval)
			select {
			case evenQueue <- data:
				evenQueue <- data
				<-timeoutChan
			case <-timeoutChan:
				<-evenQueue
			}
		}
	}()

	return evenQueue
}
