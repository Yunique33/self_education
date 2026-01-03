package main

import "fmt"

func main() {
	score := 10

	if score > 10 {
		if score > 15 {
			fmt.Println("You`re mega best")
		} else {
			fmt.Println("You`re the best!")
		}
	} else {
		fmt.Println("You lost")
	}

	if score > 15 {
		fmt.Println("You`re mega best")
	} else if score > 10 {
		fmt.Println("You`re the best!")
	} else {
		fmt.Println("You lost")
	}

	fmt.Println(score)

	switch score {
	case 12:
		fmt.Println("Dozen")
	case 21:
		fmt.Println("Blackjack")
	case 50:
		fmt.Println("You Hit")
	default:
		fmt.Println("You didn`t get")
	}

	if score < 6 || score > 16 {
		fmt.Println("You out of range!")
	}

	if score > 8 && score < 15 {
		fmt.Println("You hit it!")
	}

	subscribed := true
	if !subscribed {
		fmt.Println("Please subscribe!")
	}
}
