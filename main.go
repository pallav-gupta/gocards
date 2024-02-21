package main

import "fmt"

func main() {
	//cards := newDeck()
	//cards := newDeckFromFile("my_cards")
	//cards.saveToFile("my_cards")
	//cards.shuffle()
	//cards.print()
	//hand, remainingCards := cards.deal(5)

	//hand.print()
	//remainingCards.print()
	//alex := person{"Alex", "Anderson"}
	bob := person{
		firstName: "Bob",
		lastName:  "Max",
		contact: contactInfo{
			email:   "aaa@gmail.com",
			zipCode: 93000,
		},
	}

	//fmt.Println(alex)
	//var alex person
	//alex.firstName = "Alex"

	fmt.Println(bob)
	//fmt.Printf("%+v\n", bob)
	bob.print()

	bob.updateName("Jim")
	bob.print()
	bobPointer := &bob
	fmt.Println(bobPointer)
	bobPointer.contact.email = "abcd@gmail.com"
	bob.print()

}
