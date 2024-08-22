package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome")

	// Correct slice declaration
	var fruitList = []string{"apple", "tomato", "peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	// Append more fruits to the slice
	fruitList = append(fruitList, "mango", "banana")
	fmt.Println(fruitList)

	// Slicing the slice
	fruitList = fruitList[1:3] // Started from position 1 to 3
	fmt.Println(fruitList)

	// Creating a slice using make
	highScore := make([]int, 4)

	highScore[0] = 899
	highScore[1] = 899
	highScore[2] = 899
	highScore[3] = 899

	// Appending more scores
	highScore = append(highScore, 23, 43, 454)
	fmt.Println(highScore)

	// Sorting the slice
	sort.Ints(highScore)
	fmt.Println(highScore)

	// Corrected slice declaration
	var course = []string{"go", "ss", "ss"}
	fmt.Println(course)

	// Removing an element at index 2
	var index int = 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)
}
