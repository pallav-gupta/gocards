package main

import "fmt"

func printMap(c map[string]string) {
	for key, val := range c {
		fmt.Println("Hexcode for ", key, " is ", val)
	}
}
