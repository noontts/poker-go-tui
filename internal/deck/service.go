package deck

import (
	"math/rand"
	"time"
)

type Card struct {
	Suit string
	Rank string
}

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	var cards []Card
	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	return &Deck{cards: cards}
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *Deck) Draw() Card {
	if len(d.cards) == 0 {
		return Card{}
	}

	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}

func (d *Deck) Reset() {
	*d = *NewDeck()
}
