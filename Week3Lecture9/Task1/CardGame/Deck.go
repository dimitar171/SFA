package cardgame

import (
	"errors"
	"math/rand"
	"time"
)

func (d *Deck) New() Deck {
	var pom []Card
	for i := 1; i < 14; i++ {
		for k := 1; k < 5; k++ {
			card := Card{i, k}
			pom = append(pom, card)
		}
	}
	return Deck{pom}
}

func (d *Deck) Shuffle() Deck {

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.card), func(i, j int) { d.card[i], d.card[j] = d.card[j], d.card[i] })

	return *d

}

func (d *Deck) Deal() (*Card, error) {

	if len(d.card) == 5 {
		return nil, errors.New("this error will go to main, the deck is not Done()")
	}

	/*if len(d.card) == 0 {
		return nil, errors.New("these error will not pass to main, deck is empty")
	} */

	pom := d.card[0]
	d.card = d.card[1:]

	return &pom, nil
}
func (d *Deck) Done() bool {
	return len(d.card) == 0
}
