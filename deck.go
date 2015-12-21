package main

type newDeck []Card

func (d *newDeck) deal() {

}

// Deck a deck of card
type Deck struct {
	Cards []Card
}

// NewDeck initalize a new standard 52 card deck
func NewDeck() *Deck {
	var cards []Card

	for i := range SuitList {
		for n := range CardSet {
			card := CardSet[n]
			card.Suit = SuitList[i]
			cards = append(cards, card)
		}
	}

	return &Deck{Cards: cards}
}
