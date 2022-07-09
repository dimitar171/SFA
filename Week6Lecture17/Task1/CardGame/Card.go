package cardgame

type Cards struct {
	CardVal  int
	CardSuit int
}

var Card = []Cards{{5, 2}, {4, 2}, {11, 1}, {15, 2}, {11, 3}, {7, 2}, {13, 2}, {7, 1}, {7, 2}, {5, 1}}

func CompareCards(Card1 Cards, Card2 Cards) int {

	if Card1.CardVal < Card2.CardVal { //check values
		return -1
	}
	if Card1.CardVal > Card2.CardVal {
		return 1
	} else {
		if Card1.CardSuit < Card2.CardSuit { //check suit
			return -1
		}
		if Card1.CardSuit > Card2.CardSuit {
			return 1
		} else {
			return 0
		}
	}
}

func MaxCard(cards []Cards) Cards {
	pom := cards[0]
	var i int
	for i = 1; i < len(cards); i++ {
		s := CompareCards(pom, cards[i])
		if s == -1 {
			pom = cards[i]
		}
	}
	return pom
}
