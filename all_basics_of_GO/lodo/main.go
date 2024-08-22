package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Switch and case with a random dice roll")

	// Seed the random number generator for better unpredictability (Go versions before 1.20)
	if goVersion := runtime.Version(); goVersion < "go1.20" {
		rand.Seed(time.Now().UnixNano())
	}

	diceNumber := rand.Intn(6) + 1 // Generate a random number between 1 and 6

	fmt.Println("You rolled a", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Snake eyes! Bad luck!")
	case 2, 3:
		fmt.Println("Low roll. You might need some luck.")
	case 4, 5:
		fmt.Println("Average roll. Keep it going!")
	case 6:
		fmt.Println("Six! You rolled a perfect score!")
	default:
		fmt.Println("Unexpected dice number. Check the code!")
	}
}
