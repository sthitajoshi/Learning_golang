// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("welcome")
// 	//  var ptr *int
// 	//  fmt.Print("value of pointer ",ptr)

// 	myNumber := 23
// 	var ptr = &myNumber
// 	fmt.Println("value of actual pointer is", ptr)
// 	fmt.Println("value of actual pointer is", *ptr)

// 	*ptr = *ptr * 2
// 	fmt.Println("New value is:", myNumber)
// }

package main

import "fmt"

func main() {
	fmt.Println("welcome")

	var fruitlist [4]string
	fruitlist[0] = "aple"
	fruitlist[1] = "nb"
	fruitlist[3] = "apl"

	fmt.Println("fruits list is ", len(fruitlist))

}
