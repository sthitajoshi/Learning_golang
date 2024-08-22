package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to pizza app")

	// Create a new reader to read input from the console
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user for input
	fmt.Print("Enter a rating: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Convert the input to a float
	numRating, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Thanks for rating:", numRating)
}
