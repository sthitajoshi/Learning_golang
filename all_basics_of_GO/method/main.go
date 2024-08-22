package main

import "fmt"

func main() {
	fmt.Println("welcome")
	// there is no inheritance in golng; no super or parent

	sthita := User{"sthita ", "sthitajoshi@gmail.com", true, 14}
	fmt.Println(sthita)
	fmt.Printf("details are :%v\n", sthita)
	fmt.Printf("Name is %v and email is %v ", sthita.Name, sthita.Email)
	sthita.GetStatus()
	sthita.Newmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {

	fmt.Println("id USER ACTIVE", u.Status)
}
func (u User) Newmail() {
	u.Email = "sthita@GMAIL.com"
	fmt.Println("Email of the user is", u.Email)
}
