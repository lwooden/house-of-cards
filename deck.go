/*
Variables and functions defined in the same "pacakge" can be called freely between go files.
These strucures/methods do not have to be imported nor does their first letter have to be capitalized (marked for export)
*/
package main

import "fmt"

type deck []string

/*
This is a receiver function.
d represents an instance of the deck type and thus represents the relationship between the function and the type
Suttle resemblance of the "impl" block in Rust
Any variable of type "deck" has access to the print "function"
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	// declare two slices that hold each suit and value in a deck
	suits := []string{"Clubs", "Spades", "Hearts", "Diamond"}
	values := []string{"Ace", "One", "Two", "Three"}

	// nested for loop to iterate over both, combine a suit/value, and append to the deck list
	// _ tells go that we want to "ignore" a value
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}
