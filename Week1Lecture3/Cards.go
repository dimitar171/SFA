package main

import "fmt"

func compareCards(CardOneVal, CardOneSuit, CardTwoVal, CardTwoSuit int) int {

	if CardOneVal < CardTwoVal {
		return -1
	}
	if CardOneVal > CardTwoVal {
		return 1
	} else {
		if CardOneSuit < CardTwoSuit {
			return -1
		}
		if CardOneSuit > CardTwoSuit {
			return 1
		} else {
			return 0
		}

	}
}

func main() {
	var CardOneVal int
	var CardTwoVal int
	var CardOneSuit int
	var CardTwoSuit int

	//Enter values of cards
	fmt.Print("Enter card one suit : ")
	fmt.Scanf("%d\n", &CardOneSuit)
	fmt.Print("Enter car one value: ")
	fmt.Scanf("%d\n", &CardOneVal)
	fmt.Print("Enter card two suit : ")
	fmt.Scanf("%d\n", &CardTwoSuit)
	fmt.Print("Enter car two value: ")
	fmt.Scanf("%d", &CardTwoVal)

	res := compareCards(CardOneVal, CardOneSuit, CardTwoVal, CardTwoSuit)
	fmt.Println(res)

}
