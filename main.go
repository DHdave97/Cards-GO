package main

func main() {
	cards := newDeck()
	// cards.saveToFile("cards_dave")
	cards.suffle()
	cards.print()
}
