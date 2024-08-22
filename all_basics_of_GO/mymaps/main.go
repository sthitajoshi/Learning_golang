package main

import "fmt"

func main() {
	fmt.Println("welcomes")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println("list of langusge :", languages)
	fmt.Println("JS shorts for :", languages["js"])

	//delete(languages, "rb")
	delete(languages, "rb")
	fmt.Println("list of all lang ", languages)

	for _, value := range languages {
		fmt.Println("for key , value is ", value)
	}

}
