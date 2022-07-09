package cardgame

import (
	"math/rand"
	"testing"
	"time"
)

func TestComapare(t *testing.T) {
	//Arrange
	var expectedRes int
	var result int
	var tests = []struct {
		a, b, c, d int
	}{
		{13, 2, 10, 3},
		{9, 2, 10, 3},
		{13, 2, 13, 2},
		{13, 2, 13, 1},
		{13, 1, 13, 2},
	}
	for _, tt := range tests {
		if tt.a < tt.c { //check values
			expectedRes = -1
		}
		if tt.a > tt.c {
			expectedRes = 1
		} else {
			if tt.b < tt.d { //check suit
				expectedRes = -1
			}
			if tt.b > tt.d {
				expectedRes = 1
			} else if tt.b == tt.d {
				expectedRes = 0
			}
		}

		result = CompareCards(Cards{tt.a, tt.b}, Cards{tt.c, tt.d})
		if expectedRes != result {
			t.Errorf("Expected %d got %d ", expectedRes, result)
		}
	}

}

func TestMax(t *testing.T) {
	//Arrange

	rand.Seed(time.Now().UnixMilli())
	a, b, c, d := rand.Intn(13), rand.Intn(4), rand.Intn(13), rand.Intn(4)
	var TestCard = []Cards{{a, b}, {c, d}}
	var expectedRes Cards

	comparator := CompareCards(Cards{a, b}, Cards{c, d}) //i use the comparator function to check which card is bigger

	if comparator == 1 || comparator == 0 {
		expectedRes = Cards{a, b}
	}
	if comparator == -1 {
		expectedRes = Cards{c, d}
	}
	//Act
	result := MaxCard(TestCard)

	//Assertion
	if result != expectedRes {
		t.Errorf("Expected %d got %d ", expectedRes, result)
	}
}
