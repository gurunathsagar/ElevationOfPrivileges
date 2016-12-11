package corelogic

type GameSide struct {
	InHand      Card
	InPlay      Card
	Deck        Deck
	discardPile Card
}
