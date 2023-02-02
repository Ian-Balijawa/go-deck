package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := NewDeck()

	if len(d) != 20 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}

	if d[0] != "Ace of spades" {
		t.Errorf("Expected Ace of Spades as first card but got $v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs as first card but got $v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := NewDeck()
	deck.SaveToFile("_decktesting")

	loadedDeck := NewDeckFromFile("_decktesting")

	if len(loadedDeck) != 20 {
		t.Errorf("Expected 20 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
