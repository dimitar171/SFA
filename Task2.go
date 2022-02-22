
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

func citiesAndPrices()([]string, []int) {
	rand.Seed(time.Now().UnixMilli())
	cityChoices := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London"}
	dataPointCount := 100

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
var cities,prices = citiesAndPrices
func groupSlices(keySlice []string, valueSlice) map[string][]int {
	m := make(map[string]int)
	return  map[string]type
}