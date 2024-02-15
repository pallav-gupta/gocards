package main

import "fmt"

// Create a new type of deck
// it is a slice of strings
type deck []string

// func <receiver> print() {}
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
