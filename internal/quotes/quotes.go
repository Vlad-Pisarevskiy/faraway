package quotes

import (
	"math/rand/v2"
)

type Quoter struct {
	quotes []string
}

func NewQuoter() *Quoter {

	var quoter Quoter
	quoter.seed()

	return &quoter
}

func (q *Quoter) Random() string {

	return q.quotes[rand.IntN(len(q.quotes))]
}

func (q *Quoter) seed() {

	quotes := []string{
		"The only true wisdom is in knowing you know nothing.",
		"He who knows that enough is enough will always have enough.",
		"The obstacle is the way.",
		"We suffer more often in imagination than in reality.",
		"A journey of a thousand miles begins with a single step.",
		"What we do now echoes in eternity.",
		"The unexamined life is not worth living.",
		"Waste no more time arguing what a good man should be. Be one.",
		"Knowing yourself is the beginning of all wisdom.",
		"It does not matter how slowly you go as long as you do not stop.",
	}

	q.quotes = quotes
}
