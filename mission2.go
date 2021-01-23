package main

import (
	"fmt"
	"strings"
	"time"
)

func m2Briefing() {
	fmt.Println("\nCurrent Location: CIA camp in Hasaka")
	time.Sleep(4 * time.Second)
	fmt.Println("\nNext day you are assigned a new mission, you are attacking ISIS territory in Hasaka!")
	time.Sleep(4 * time.Second)
	fmt.Println("\nSRGT Dave wakes you up and you get all your gear on, and eat breakfast. 2 hours later you will leave in the HMV.")
	m2ConfrontationOne()
}

func m2ConfrontationOne() {
	var choice string

	fmt.Println("\nSRGT Dave screams at you once again to get in the HMV, you already know the deal and you enter cautiously")
	time.Sleep(4 * time.Second)
	fmt.Println("\nSome time later, you begin to drive into battle")
	time.Sleep(4 * time.Second)
	fmt.Print("\nYou have to drive through Hasaka, and flank them do u wish to proceed? (y/n): ")
	fmt.Scan(&choice)

	choiceL := strings.ToLower(choice)

	if choiceL == "n" {
		fmt.Print("\nMission failed, wrong choice")
		m2Briefing()
	} else if choiceL == "y" {
		fmt.Println("\nYou continue to flank them, exit the HMV, and begin to fire upon the enemies")
		time.Sleep(4 * time.Second)
	}
	fmt.Println("\nYou get hit across the back of the head and blackout as the enemies stop firing back at you")
	time.Sleep(4 * time.Second)
	fmt.Println("\nYou wake up some time later, to see SRGT Dave above you with a pistol to your head. Sorry it had to end this way, you were just a liability. ISIS forever. He says as he shoots you in the head. Game over, you are dead.")
}
