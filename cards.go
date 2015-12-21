package main

// Card a playing card
type Card struct {
	Suit     Suit
	Value    int
	Type     CardType
	FaceType FaceType
}

// Greater compares two cards if this card it greater returns true
func (c *Card) Greater(cc *Card) bool {
	if c.Value > cc.Value {
		return true
	} else if (c.Value == c.Value) && (c.Type == c.Type) {
		if (c.Type == FaceCard) && (cc.Type == FaceCard) {
			switch c.FaceType {
			case Jack:
				if (cc.FaceType == Queen) || (cc.FaceType == King) {
					return false
				}
				return true
			case Queen:
				if cc.FaceType == King {
					return false
				}
				return true
			case King:
				if cc.FaceType == King {
					return false
				}
				return true
			}
		} else if (c.Type == AceCard) && (cc.Type != AceCard) {
			return true
		} else if (c.Type == FaceCard) && (cc.Type == AceCard) {
			return false
		}
	}
	return false
}

func (c *Card) Equal(cc *Card) bool {
	return false
}

var two = Card{Value: 2, Type: NumberCard}
var three = Card{Value: 3, Type: NumberCard}
var four = Card{Value: 4, Type: NumberCard}
var five = Card{Value: 5, Type: NumberCard}
var six = Card{Value: 6, Type: NumberCard}
var seven = Card{Value: 7, Type: NumberCard}
var eight = Card{Value: 8, Type: NumberCard}
var nine = Card{Value: 9, Type: NumberCard}
var ten = Card{Value: 10, Type: NumberCard}
var jack = Card{Value: 10, Type: FaceCard, FaceType: Jack}
var king = Card{Value: 10, Type: FaceCard, FaceType: King}
var queen = Card{Value: 10, Type: FaceCard, FaceType: Queen}
var ace = Card{Value: 1, Type: AceCard}
var joker = Card{Value: 0, Type: JokerCard}

// CardSet a set of cards for a specific suit
// by default the suit is not set and the joker is not included
var CardSet = []Card{two, three, four, five, six, seven, eight, nine, ten, jack, queen, king, ace}

//go:generate stringer -type=CardType

// CardType the type of card
type CardType int

const (
	_ CardType = iota
	// NumberCard a number card type
	NumberCard
	// FaceCard a face card type
	FaceCard
	// AceCard an ace card type
	AceCard
	// JokerCard a joker card type
	JokerCard
)

//go:generate stringer -type=FaceType

// FaceType type of face card
type FaceType int

const (
	_ FaceType = iota
	// Jack jack face card
	Jack
	// Queen queen face card
	Queen
	// King face card
	King
)
