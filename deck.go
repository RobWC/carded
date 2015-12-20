package main

type newDeck []Card

func (d *newDeck) deal() {

}

// Deck a deck of cards
type Deck struct {
	Cards []Card
}

// NewDeck initalize a new standard 52 card deck
func NewDeck() *Deck {
	var cards []Card

	for i := range SuitList {
		for n := range CardSet {
			CardSet[n].Suit = SuitList[i]
			cards = append(cards, CardSet[n])
		}
	}

	return &Deck{Cards: cards}
}
