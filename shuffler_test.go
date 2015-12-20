package main

import "testing"

func TestFoldingShuffler(t *testing.T) {
	d := NewDeck()
	FoldingShuffler(d)
	if len(d.Cards) != 52 {
		t.Fatalf("Too many cards in the deck %d", len(d.Cards))
	}
	for i := range d.Cards {
		t.Log("card", d.Cards[i])
	}
}
