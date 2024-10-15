package main

import "fmt"

func GetName() string {
	name := ""

	fmt.Println("Welcome to Tim's Casino...")
	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s, lets's play!\n", name)
	return name
}

func GetBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("Enter your bet, or 0 to quit (balance = $%d): ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Bet cannot be larger than balance.")
		} else {
			break
		}
	}
	return bet
}
