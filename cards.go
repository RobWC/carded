package main

type Card struct {
	Suit  Suit
	Value int
	Type  CardType
}

var Two = Card{Value: 2, Type: NumberCard}
var Three = Card{Value: 3, Type: NumberCard}
var Four = Card{Value: 4, Type: NumberCard}
var Five = Card{Value: 5, Type: NumberCard}
var Six = Card{Value: 6, Type: NumberCard}
var Seven = Card{Value: 7, Type: NumberCard}
var Eight = Card{Value: 8, Type: NumberCard}
var Nine = Card{Value: 9, Type: NumberCard}
var Ten = Card{Value: 10, Type: NumberCard}
var Jack = Card{Value: 10, Type: FaceCard}
var King = Card{Value: 10, Type: FaceCard}
var Queen = Card{Value: 10, Type: FaceCard}
var Ace = Card{Value: 1, Type: AceCard}
var Joker = Card{Value: 0, Type: JokerCard}

var CardSet []Card = []Card{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

//go:generate stringer -type=CardType

type CardType int

const (
	NumberCard CardType = iota
	FaceCard
	AceCard
	JokerCard
)
