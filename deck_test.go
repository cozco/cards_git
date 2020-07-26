package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 28 {
		//t.Errorf("Expected deck length of 28, but got", len(d))
		t.Errorf("Expected deck length of 28, but got %v", len(d))
	}

	if d[0] != "King of Spades" {
		t.Errorf("Expected King of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Diamonds" {
		t.Errorf("Expected Four of Diamonds, but got %v", d[len(d)-1])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	d = newDeckFromFile("_decktesting")
	if len(d) != 28 {
		t.Errorf("Expected deck length of 28, but got %v", len(d))
	}
	os.Remove("_decktesting")
}
