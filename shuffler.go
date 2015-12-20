package main

import (
	"math/rand"
	"time"
)

//Shuffler interface for different types of card shufflers
// each shuffler will change how it decides to randomize the deck
// by using a random shuffler you will increase the entropy and
type Shuffler interface {
	Shuffle(*Deck)
}

// FoldingShuffler shuffles cards by splitting the deck and folding them together
func FoldingShuffler(d *Deck) {
	var neworder []Card

	for i := 0; i < len(d.Cards); i = i + 1 {
		neworder = append(neworder, d.Cards[i])
	}

	d.Cards = neworder
}

// RandShuffler shuffle cards in a random order
func RandShuffler(d *Deck) {
	rand.Seed(int64(time.Now().Nanosecond()))
}
