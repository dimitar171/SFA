package carddraw

import (
	"fmt"

	cardgame "github.com/dimitar171/SFA/Week3Lecture8/Task1/CardGame"
)

type dealer interface {
	Deal() []cardgame.Card
}

func DrawAllCards(d dealer) []cardgame.Card {
	card := d.Deal()
	for i := 0; i < 52; i++ { //removing items from the deck
		fmt.Println(card)
		card = d.Deal()
	}

	return nil
}
