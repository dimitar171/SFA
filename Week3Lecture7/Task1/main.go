package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	deck []Card
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
	rand.Shuffle(len(d.deck), func(i, j int) { d.deck[i], d.deck[j] = d.deck[j], d.deck[i] })

	return *d

}

func (d *Deck) Deal() Deck {
	if len(d.deck) == 0 {
		d.deck = nil
		return *d
	}
	d.deck = d.deck[1:]
	return *d
}

func main() {
	deck := Deck{}
	fmt.Println(deck) //this checks an empty deck

	deck = deck.New() //filing the deck
	fmt.Println(deck)

	deck = deck.Shuffle() //shuffling it
	fmt.Println(deck)

	for i := 0; i < 30; i++ { //removing items from the deck
		deck = deck.Deal()
	}
	fmt.Println(deck)
}
