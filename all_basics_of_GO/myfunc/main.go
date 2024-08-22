package main

import "fmt"

func main() {
	greater()
	fmt.Println("welcome")
}

func greater() {
	fmt.Println("namste")

	result := adder(3, 5)
	fmt.Println("reslut is", result)

	proRes := proadder(2, 3, 5, 5)
	fmt.Println("Pro result is:", proRes)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proadder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
