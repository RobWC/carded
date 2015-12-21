package main

import "testing"

func TestShuffler(t *testing.T) {
	d := NewDeck()
	Shuffler(d)
	/*	if len(d.Cards) != 52 {
		t.Fatalf("Too many cards in the deck %d", len(d.Cards))
	}*/
	t.Log("Total Cards:", len(d.Cards))
	for i := range d.Cards {
		t.Log("card", d.Cards[i])
	}
}
