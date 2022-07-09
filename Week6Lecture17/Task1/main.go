package main

import (
	"fmt"

	cardgame "github.com/dimitar171/SFA/Week6Lecture17/Task1/CardGame"
)

type Cards struct {
	CardVal  int
	CardSuit int
}

func main() {

	res2 := cardgame.MaxCard(cardgame.Card)
	//prints the strongest card Val and Suit
	fmt.Println(res2.CardVal, res2.CardSuit)
}
