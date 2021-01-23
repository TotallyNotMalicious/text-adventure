package main

import (
	"fmt"
	"strings"
	"time"
)

func m1Briefing(name string) {
	var option string

	fmt.Println("\nCurrent Task: you must infiltrate Basra Afghanistan, and free the civilians being held as hostage.\n\nCurrent Location: CIA camp in Herat")
	time.Sleep(7 * time.Second)
	fmt.Println("\nYou have options on what classes of weapons you desire to have.")
	time.Sleep(5 * time.Second)
	fmt.Println("\nOption 1: M16, with the following attachments. holo sight, foregrip, and 120 rounds of 5.56 ammo. Secondary weapon of Glock 18 with 2 clips of .45 ACP")
	time.Sleep(7 * time.Second)
	fmt.Println("\nOption 2: LMG M60, with 400 rounds of 7.62 ammo. Secondary weapon of 1911 Baretta, with 1 clip of 9mm")
	fmt.Print("\nOption number: ")
	time.Sleep(4 * time.Second)
	fmt.Scan(&option)

	fmt.Printf("\nYou chose option: %s", option)
	m1ConfrontationOne(name)
	time.Sleep(5 * time.Second)
}

func m1ConfrontationOne(name string) {
	var choice string

	fmt.Printf("\n\nSRGT Dave yells out your name. HEY %s get in the HMV now we got to go to Basra, hurry maggot!", name)
	time.Sleep(4 * time.Second)
	fmt.Print("\nDo you listen to SRGT Dave? (y/n): ")
	fmt.Scan(&choice)

	choiceL := strings.ToLower(choice)

	if choiceL == "y" {
		m1ConfrontationTwo()
	} else {
		fmt.Print("\nMission failed, game over")
		main()
	}
}

func m1ConfrontationTwo() {
	var choice string
	var choice2 string
	var choice3 string

	fmt.Println("\nYou are on your way to Basra in the HMV")
	time.Sleep(4 * time.Second)
	fmt.Println("\nALRIGHT MAGGOT GET OUT NOW! SRGT Dave says.")
	time.Sleep(4 * time.Second)
	fmt.Println("\nSomething is off, it's too silent in Basra, as there is constant gunfire.")
	time.Sleep(6 * time.Second)
	fmt.Println("\nTHE HMV EXPLODES FROM AN ISIS RPG ROCKET RIGHT BEHIND YOU AS YOU QUICKLY GET TO COVER!")
	time.Sleep(4 * time.Second)
	fmt.Println("\nThere are bullets flying everywhere, you gauge on how many enemies there are.")
	time.Sleep(4 * time.Second)
	fmt.Print("\nThere is an enemy on top of left compound, do you take the shot and risk getting hit? (y/n): ")
	fmt.Scan(&choice)

	choiceL := strings.ToLower(choice)

	if choiceL == "n" {
		fmt.Print("\nMission failed, game over. You got shot from an enemy to the side of you.")
		main()
	} else {
		fmt.Println("\nYou successfully kill the enemy on the left compound!")
	}
	time.Sleep(4 * time.Second)
	fmt.Println("\nSRGT Dave clears out the rest of the enemies, your team moves closer to Basra! There is more fire coming from on top of the compounds! Get to cover immediately.")
	time.Sleep(4 * time.Second)
	fmt.Print("\nDo you take the shot on the enemy on the right roof? (y/n): ")
	fmt.Scan(&choice2)

	choice2L := strings.ToLower(choice2)

	if choice2L == "n" {
		fmt.Println("\nYou get hit in the leg and reactively you fire a few shots inaccurately, current health: 85")
	} else if choice2L == "y" {
		fmt.Println("\nYou take the shot as you kill the enemy on the right roof")
	}
	fmt.Print("\nWould you like to reload? (y/n): ")
	fmt.Scan(&choice3)

	choice3L := strings.ToLower(choice3)

	if choice3L == "n" {
		fmt.Print("\nYou are dead, game over")
		main()
	} else if choice3L == "y" {
		fmt.Println("\nYou reload your gun and kill the remaining person on left roof")
	}
	m1ConfrontationThree()
}

func m1ConfrontationThree() {
	var choice string
	var choice2 string
	var choice3 string

	fmt.Println("\nYour team moves up and you have sight of the compound where the hostages are being stored")
	time.Sleep(4 * time.Second)
	fmt.Print("\nYou enter the stairs of the compound, but you gather intel that there are 4 enemies in the compound. Do you enter the compound and risk getting shot? (y/n): ")
	fmt.Scan(&choice)

	choiceL := strings.ToLower(choice)

	if choiceL == "n" {
		fmt.Print("\nYou get sniped from a nearby building as your decision is hesitant")
		main()
	} else if choiceL == "y" {
		fmt.Println("\nYou decide to enter the compound")
	}
	fmt.Print("\nSRGT Dave is behind you and your team has your six! You are the first to enter the compound, there is an enemy in the Kitchen and he hears you breaking in, and rushes you, you have a shot. Do you take it? (y/n): ")
	fmt.Scan(&choice2)

	choice2L := strings.ToLower(choice2)

	if choice2L == "n" {
		fmt.Print("\nThe enemy kills you, game over")
		main()
	} else if choice2L == "y" {
		fmt.Println("\nYou take the shot and injure him as he falls to the ground")
		time.Sleep(4 * time.Second)
	}
	fmt.Println("\nThere is another enemy in the kids bedroom, you begin to notice as you approach the enemy, take the shot and kill him")
	fmt.Print("\nYou are out of ammo, you need to reload, do you? (y/n): ")
	fmt.Scan(&choice3)

	choice3L := strings.ToLower(choice3)

	if choice3L == "n" {
		fmt.Println("\nYou continue without reloading")
	} else if choice3L == "y" {
		fmt.Println("\nYou reload and continue")
	}
	time.Sleep(4 * time.Second)
	fmt.Println("\nThere are two enemies in the basement, you move down the basement with your team on your back! You kill 1, but get hit. Current health: 50")
	time.Sleep(4 * time.Second)
	fmt.Println("\nMission succsessful, the new HMV is enroute to your location, congrats you saved the hostages!")
	time.Sleep(4 * time.Second)
	m2Briefing()
}
