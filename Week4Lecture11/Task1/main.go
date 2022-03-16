package main

import (
	"log"
	"sync"
)

func processEven(inputs []int) chan int {
	evenQueue := make(chan int, len(inputs)) //buffered
	outChan := make(chan int)                //unbafered
	var wg sync.WaitGroup
	go func() {

		for _, inputToProcess := range inputs { //worker
			if inputToProcess%2 == 0 { //check for even num
				wg.Add(1)
				evenQueue <- inputToProcess

				go func(input int) {
					<-evenQueue //emptying the buffer
					outChan <- input
					wg.Done()

				}(inputToProcess)
			}
		}
		wg.Wait()
		close(outChan)
	}()
	log.Println("Chanel Even send") //chanel is send before the numbers
	return outChan
}

func processOdd(inputs []int) chan int {
	oddQueue := make(chan int, len(inputs)) //buffered
	outChan := make(chan int)               //unbafered
	var wg sync.WaitGroup
	go func() {

		for _, inputToProcess := range inputs { //worker
			if inputToProcess%2 != 0 { //check for odd num
				wg.Add(1)
				oddQueue <- inputToProcess

				go func(input int) {
					<-oddQueue //emptying the buffer
					outChan <- input
					wg.Done()

				}(inputToProcess)
			}
		}
		wg.Wait()
		close(outChan)
	}()
	log.Println("Chanel Odd send") //chanel is send before the numbers
	return outChan
}

func main() {
	inputs := []int{1, 17, 34, 56, 2, 8, 15, 30, 22, 23}

	evenCh := processEven(inputs)
	oddCh := processOdd(inputs)
	for inp := range evenCh {
		log.Println("Even Num:", inp)
	}
	for inp := range oddCh {
		log.Println("Odd Num:", inp)
	}

}
