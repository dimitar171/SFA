package main

import (
	"fmt"
	"sync"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
}

func (cp *ConcurrentPrinter) printFooBar(times int) {

	for i := 0; i < times; i++ {
		cp.Add(1)
		go func() {
			fmt.Print("foo")
			cp.Done()
		}()
		cp.Wait()
		cp.Add(1)
		go func() {
			fmt.Print("bar")
			cp.Done()
		}()

	}
}

func main() {
	times := 100
	cp := &ConcurrentPrinter{}
	cp.printFooBar(times)
	cp.Wait()
}
