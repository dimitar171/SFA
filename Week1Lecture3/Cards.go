package main

import (
	"errors"
	"fmt"
)

func compareCards(CardOneVal, CardOneSuit, CardTwoVal, CardTwoSuit int) (int, error) {
	//error handling

	if CardOneVal < 2 || CardOneVal > 13 || CardTwoVal < 2 || CardTwoVal > 13 {
		return -1, errors.New("Card value should be between 2 and 13, try again")
	}
	if CardOneSuit < 1 || CardOneSuit > 4 || CardTwoSuit < 1 || CardTwoSuit > 4 {
		return -1, errors.New("Card suit should be between 1 and 4, try again")
	}

	//function logic flow

	if CardOneVal < CardTwoVal { //check values
		return -1, nil
	}
	if CardOneVal > CardTwoVal {
		return 1, nil
	} else {
		if CardOneSuit < CardTwoSuit { //check suit
			return -1, nil
		}
		if CardOneSuit > CardTwoSuit {
			return 1, nil
		} else {
			return 0, nil
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

	res, err := compareCards(CardOneVal, CardOneSuit, CardTwoVal, CardTwoSuit)
	if err != nil {
		// Handle the error
		fmt.Println(err)
	} else {
		// No errors
		fmt.Println(res)
	}

}
