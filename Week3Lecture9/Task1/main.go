package main

import (
	"fmt"
	"log"

	carddraw "github.com/dimitar171/SFA/Week3Lecture9/Task1/CardDraw"
	cardgame "github.com/dimitar171/SFA/Week3Lecture9/Task1/CardGame"
)

func main() {
	deck := cardgame.Deck{}
	deck = deck.New()
	fmt.Println(deck)

	res, er := carddraw.DrawAllCards(&deck)

	if er != nil {
		log.Fatal(er) //prints the error
	}

	fmt.Println(res) //prints last card or empty card

}
