package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println(primesAndSleep(100, 1))
	fmt.Println(goPrimesAndSleep(100, 1))
}

func goPrimesAndSleep(n int, sleep time.Duration) []int {
	res := []int{}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for k := 2; k < n; k++ {

			for i := 2; i < n; i++ {
				if k%i == 0 {
					if k == i {
						res = append(res, k)
						time.Sleep(sleep)
					}
					break
				}

			}

		}
	}()

	wg.Wait()
	return res
}

func primesAndSleep(n int, sleep time.Duration) []int {
	res := []int{}
	for k := 2; k < n; k++ {
		for i := 2; i < n; i++ {
			if k%i == 0 {
				time.Sleep(sleep)
				if k == i {
					res = append(res, k)
				}
				break
			}
		}
	}
	return res
}
