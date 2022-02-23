package main

import (
	"fmt"
	"math/rand"
	"time"
)

func citiesAndPrices() ([]string, []int) {
	rand.Seed(time.Now().UnixMilli())
	cityChoices := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London"}
	dataPointCount := 10

	// randomly choose cities
	cities := make([]string, dataPointCount)
	for i := range cities {
		cities[i] = cityChoices[rand.Intn(len(cityChoices))]
	}

	prices := make([]int, dataPointCount)
	for i := range prices {
		prices[i] = rand.Intn(100)
	}

	return cities, prices
}

func groupSlices(keySlice []string, valueSlice []int) map[string][]int {

	m := make(map[string][]int)

	for _, v := range keySlice {
		m[v] = valueSlice
	}
	return m
}

func main() {
	a, b := citiesAndPrices()

	var res = groupSlices(a, b)

	for k, v := range res {
		fmt.Printf("key[%s] value[%d]\n", k, v)
	}

}
