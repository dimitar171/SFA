package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	card []Card
}

type Card struct {
	cardVal  int
	cardSuit int
}

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

func (d *Deck) Deal() *Card {

	if len(d.card) == 0 {
		return nil
	}
	pom := &d.card[0]
	d.card = d.card[1:]

	return pom
}

func main() {
	deck := Deck{}
	deck = deck.New() //filing the deck
	fmt.Println(deck)

	card := deck.Deal()
	fmt.Println(card)

	for i := 0; i < 5; i++ { //removing items from the deck
		card = deck.Deal()
		fmt.Println(*card)
	}
	fmt.Println(deck)
}
