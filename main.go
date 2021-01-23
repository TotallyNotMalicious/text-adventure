package main

import (
	"fmt"
)

func main() {
	var firstName string
	var lastName string

	fmt.Print("\nWhat is your characters first name? ")
	fmt.Scan(&firstName)
	fmt.Print("What is your characters last name? ")
	fmt.Scan(&lastName)
	fmt.Printf("\nWelcome, %s %s!\n", firstName, lastName)
	m1Briefing(firstName)
}
