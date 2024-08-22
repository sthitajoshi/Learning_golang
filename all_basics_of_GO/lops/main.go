package main

import "fmt"

func main() {
	fmt.Println("wel")

	days := []string{"sunday", "monday", "tuesday", "wednesday"}

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }
	// for i:= range days{
	// 	fmt.Println(days[i])
	// }
	for _, day := range days {
		fmt.Printf("index is__ and value is %v\n", day)
	}
}
