package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://localhost:3000/"

func main() {
	fmt.Println("welcome")

	Response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Responce is of type :%T", Response)

	Response.Body.Close()

	databytes, err := ioutil.ReadAll(Response.Body)
	if err != nil {
		panic(err)
	}
	constant := string(databytes)
	fmt.Println(constant)
}
