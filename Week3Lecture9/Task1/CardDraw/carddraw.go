package carddraw

import (
	"fmt"

	cardgame "github.com/dimitar171/SFA/Week3Lecture9/Task1/CardGame"
)

type dealer interface {
	Deal() (*cardgame.Card, error)

	Done() bool
}

func DrawAllCards(d dealer) (*cardgame.Card, error) {
	card, er := d.Deal() //assign the deck

	for i := 0; i < 53; i++ {
		if er != nil { //checks for error
			if d.Done() { //checks if Deck is dealt

				return card, nil
			}
			return nil, er //if its has an error, but its not Done - we pass the error to main
		}

		fmt.Println(*card)  //prints every dealt card
		card, er = d.Deal() //asigns new deck without the dealt card
	}

	return card, er
}
