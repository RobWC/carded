package main

//go:generate stringer -type=Suit

// Suit card suits
type Suit int

const (
	_ Suit = iota
	// Heart suit
	Heart
	// Diamond suit
	Diamond
	// Spade  suit
	Spade
	// Club suit
	Club
)

// SuitList list for all availble suits
var SuitList = []Suit{Heart, Diamond, Spade, Club}
