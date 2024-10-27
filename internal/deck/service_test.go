package deck

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	if len(deck.cards) != 52 {
		t.Errorf("Expected 52 cards, got %d", len(deck.cards))
	}
}

func TestShuffle(t *testing.T) {
	deck := NewDeck()
	firstCard := deck.cards[0]
	deck.Shuffle()
	if deck.cards[0] == firstCard {
		t.Errorf("Expected first card to be different after shuffle")
	}
}

func TestDraw(t *testing.T) {
	deck := NewDeck()
	card := deck.Draw()
	if len(deck.cards) != 51 {
		t.Errorf("Expected 51 cards after draw, got %d", len(deck.cards))
	}
	if card == (Card{}) {
		t.Errorf("Expected a valid card, got an empty card")
	}
}

func TestReset(t *testing.T) {
	deck := NewDeck()
	deck.Draw()
	deck.Reset()
	if len(deck.cards) != 52 {
		t.Errorf("Expected 52 cards after reset, got %d", len(deck.cards))
	}
}
