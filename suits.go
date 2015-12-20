package main

//go:generate stringer -type=Suit

// Suit card suits
type Suit int

const (
	// Heart suit
	Heart Suit = iota
	// Diamond suit
	Diamond
	// Spade  suit
	Spade
	// Club suit
	Club
)

var SuitList []Suit = []Suit{Heart, Diamond, Spade, Club}
