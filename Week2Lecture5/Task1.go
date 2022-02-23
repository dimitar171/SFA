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
	if CardOneSuit < 0 || CardOneSuit > 3 || CardTwoSuit < 0 || CardTwoSuit > 3 {
		return -1, errors.New("Card suit should be club,diamond,heart or spade, try again")
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

	type person struct {
		name string
		age  int
	}
	type CardSuit = int
	const (
		club CardSuit = iota
		diamond
		heart
		spade
	)

	//Enter values of cards
	var CardOneVal = []int{7, 5, 3, 2, 14, 7, 5, 1, 12, 8}
	var CardTwoVal = []int{7, 3, 3, 4, 8, 7, 5, 6, 12, 8}

	var CardOneSuit = []CardSuit{diamond, heart, heart, diamond, heart, diamond, heart, spade, spade, heart}
	var CardTwoSuit = []CardSuit{diamond, heart, spade, spade, diamond, diamond, spade, diamond, 5, diamond}

	var i int
	for i = 0; i < 10; i++ {

		res, err := compareCards(CardOneVal[i], CardOneSuit[i], CardTwoVal[i], CardTwoSuit[i])
		if err != nil {
			// Handle the error
			fmt.Println(err)
		} else {
			// No errors
			fmt.Println(res)
		}
	}
}
