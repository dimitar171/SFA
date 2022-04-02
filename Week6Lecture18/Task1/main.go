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
	for k := 2; k < n; k++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			for i := 2; i < n; i++ {
				if k%i == 0 {
					time.Sleep(sleep)
					if k == i {
						res = append(res, k)
					}
					break
				}

			}

		}(k)
	}
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
