package main

import "fmt"

func main() {

	var month = []int{2, 15, 3, 1, 2, 7, 2, 1, 12, 0}
	var year = []int{2004, 1973, 1392, 1994, 2004, 1117, 1995, 2004, 1222, 1998}
	var i int
	for i = 0; i < 10; i++ {

		res, err := daysInMonth(month[i], year[i])
		if err == false {
			fmt.Println("Enter month between 1 and 12")
		} else {
			fmt.Println(res)
		}
	}
}

func daysInMonth(month int, year int) (int, bool) {

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31, true
	case 4, 6, 9, 11:
		return 30, true
	case 2:
		switch {
		case year%4 != 0:
			return 29, true
		case year%4 == 0:
			return 28, true
		}
	}
	return -1, false
}
