package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := NewDeck()
	if len(d.Cards) != 52 {
		t.Fatalf("Wrong number of cards %d", len(d.Cards))
	}
}

func TestNewDeckOrder(t *testing.T) {
	d := NewDeck()
	for i := range d.Cards {
		t.Log(d.Cards[i])
	}
}
