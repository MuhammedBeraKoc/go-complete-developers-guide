package main

func main() {
	cards := newDeckFromFile("myDeck.txt")
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}