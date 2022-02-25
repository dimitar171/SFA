package main

import (
	"errors"
	"fmt"
)

type Cards struct {
	CardVal  int
	CardSuit int
}

func compareCards(Card1 Cards, Card2 Cards) (int, error) {
	//error handling

	if Card1.CardVal < 2 || Card1.CardVal > 13 || Card2.CardVal < 2 || Card2.CardVal > 13 {
		return -1, errors.New("Card value should be between 2 and 13, try again")
	}
	if Card1.CardSuit < 0 || Card1.CardSuit > 3 || Card2.CardSuit < 0 || Card2.CardSuit > 3 {
		return -1, errors.New("Card suit should be club,diamond,heart or spade, try again")
	}
	//function logic flow

	if Card1.CardVal < Card2.CardVal { //check values
		return -1, nil
	}
	if Card1.CardVal > Card2.CardVal {
		return 1, nil
	} else {
		if Card1.CardSuit < Card2.CardSuit { //check suit
			return -1, nil
		}
		if Card1.CardSuit > Card2.CardSuit {
			return 1, nil
		} else {
			return 0, nil
		}

	}
}

func maxCard(cards []Cards) Cards {
	pom := cards[0]
	var i int
	for i = 1; i < 10; i++ {
		s, err := compareCards(pom, cards[i])
		if err != nil || s == -1 {
			pom = cards[i]
		}
	}
	return pom
}

func main() {

	var Card = []Cards{{5, 2}, {4, 2}, {11, 1}, {15, 2}, {11, 3}, {7, 2}, {13, 2}, {7, 1}, {7, 2}, {5, 1}}

	res2 := maxCard(Card)
	//prints the strongest card Val and Suit
	fmt.Println(res2.CardVal, res2.CardSuit)
}
