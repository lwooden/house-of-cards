package main

// func newCard(card string) string {
// 	return card
// }

func main() {

	cards := newDeck()
	cards.print()

	// card := "Ace of Spades" // shorthand syntax

	// card := newCard("King of Diamonds")
	// deck := []string{"2 of Hearts", "Jack of Diamonds"} // slice
	// deck = append(deck, "8 of Hearts")                  // does not modify the existing slice, but returns a new slice and assigns to variable
	// cards := deck{"2 of Hearts", "Jack of Diamonds"}
	// fmt.Println("The Deck: ")

	// iterate over the deck and return each card; ignore the index value (don't need it)
	// for _, c := range cards {
	// 	fmt.Println(c)
	// }

	// fmt.Println("A Card: ", card)
	// fmt.Println("The Deck: ", deck)

}
