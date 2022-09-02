package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Got unexpected first card: %v", d[0])
	}
	if d[len(d) - 1] != "King of Clubs" {
		t.Errorf("Got unexpected last card: %v", d[0])
	}
}

func TestSaveToFileNewDeckFromFile(t *testing.T) {
	testingFilename := "_decktesting"

	os.Remove(testingFilename)

	deck := newDeck()
	deck.saveToFile(testingFilename)

	newDeck := newDeckFromFile(testingFilename)

	deckLength := len(newDeck)
	if deckLength != 52 {
		t.Errorf("Expected 52 cards in deck, but got %v", deckLength)
	}
	os.Remove(testingFilename)
}