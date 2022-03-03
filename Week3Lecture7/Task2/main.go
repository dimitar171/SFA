package main

import (
	"fmt"
)

type Cards struct {
	CardVal  int
	CardSuit int
}

type CardComparator func(Card1 Cards, Card2 Cards) int

func compareCards(Card1 Cards, Card2 Cards) int {
	if Card1.CardVal < Card2.CardVal {
		return -1
	}
	if Card1.CardVal > Card2.CardVal {
		return 1
	} else {
		if Card1.CardSuit < Card2.CardSuit {
			return -1
		}
		if Card1.CardSuit > Card2.CardSuit {
			return 1
		} else {
			return 0
		}
	}
}

//func with ref as a paramater
func maxCard(cards []Cards, comparatorFunc CardComparator) Cards {
	pom := cards[0]
	for i := 1; i < 10; i++ {
		s := comparatorFunc(pom, cards[i])
		if s == -1 {
			pom = cards[i]
		}
	}
	return pom
}

//func with anonymous as a param
func maxA(cards []Cards, compFunc func(Card1 Cards, Card2 Cards) int) Cards {
	pom := cards[0]
	for i := 1; i < 10; i++ {
		s := compFunc(pom, cards[i])
		if s == -1 {
			pom = cards[i]
		}
	}
	return pom
}

func main() {
	//i removed error handling to make the code more readable
	var Card = []Cards{{5, 2}, {4, 2}, {11, 1}, {11, 2}, {11, 3}, {7, 2}, {13, 2}, {7, 1}, {7, 2}, {5, 1}}

	//passing as an anonymous function
	resA := maxA(Card,
		func(Card1 Cards, Card2 Cards) int {

			if Card1.CardVal < Card2.CardVal {
				return -1
			}
			if Card1.CardVal > Card2.CardVal {
				return 1
			} else {
				if Card1.CardSuit < Card2.CardSuit {
					return -1
				}
				if Card1.CardSuit > Card2.CardSuit {
					return 1
				} else {
					return 0
				}
			}
		})
	fmt.Println(resA)

	//passing as a reference
	resB := maxCard(Card, compareCards)
	fmt.Println(resB)

}
