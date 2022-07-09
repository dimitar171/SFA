package main

import (
	"fmt"

	carddraw "github.com/dimitar171/SFA/Week3Lecture8/Task1/CardDraw"
	cardgame "github.com/dimitar171/SFA/Week3Lecture8/Task1/CardGame"
)

func main() {
	deck := cardgame.Deck{}
	deck = deck.New()
	fmt.Println(deck)
	carddraw.DrawAllCards(&deck)
	fmt.Println(deck)

}
