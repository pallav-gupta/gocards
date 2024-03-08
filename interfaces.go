package main

import "fmt"

type bot interface {
	getGreeting() string
}

type engBot struct{}
type spanishBot struct{}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (engBot) getGreeting() string {
	return "hello!!!"
}

func (spanishBot) getGreeting() string {
	return "hola !!!"
}
