package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("fuck")
	defer fmt.Println("you")
	defer fmt.Println("wlow")
	fmt.Println("hello")
	myDefer()
	// last in first our
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
