package main

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

	/*
		bob := person{
			firstName: "Bob",
			lastName:  "Max",
			contact: contactInfo{
				email:   "aaa@gmail.com",
				zipCode: 93000,
			},
		}
	*/

	//fmt.Println(alex)
	//var alex person
	//alex.firstName = "Alex"

	//fmt.Println(bob)
	//fmt.Printf("%+v\n", bob)
	//bob.print()

	//bob.updateName("Jim")
	//bob.print()

	// address to value *name
	// value to address &name
	//bobPointer := &bob
	//fmt.Println("Pointer: ", bobPointer)
	//bobPointer.updateName("jimmy")

	//bobPointer.contact.email = "abcd@gmail.com"
	//bob.print()
	/*
		colors := map[string]string{
			"red":   "#ff0000",
			"green": "#4bf745",
			"white": "#ffffff",
		}

		//var colors map[string]string

		//colors := make(map[string]string)

		//colors["white"] = "#ffffff"
		//delete(colors, "white")

		fmt.Println(colors)
		printMap(colors)
		addItem(colors, "blue", "xxxxx1")
		printMap(colors)
	*/
	/*
		eb := engBot{}
		sb := spanishBot{}
		printGreeting(eb)
		printGreeting(sb)
	*/

	//httpGet()

	/*
		tr := traingle{base: 10, height: 5}
		sq := squire{sideLenth: 2.5}
		printArea(tr)
		printArea(sq)
	*/

	/*
		args := os.Args[1:]
		f, _ := os.Open(args[0])

		io.Copy(os.Stdout, f)
	*/

	links := getUrls("urls.txt")
	checkLinks(links)

}
