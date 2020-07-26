package main

func newCard() string {
	return "Five of Diamonds"
}

func main() {

	cards := newDeck()
	cards.shuffle()
	cards.PrintDeck()

	//hand, remainingCards := deal(cards, 5)

	//hand.PrintDeck()
	//remainingCards.PrintDeck()

	//cards.PrintDeck()

}
