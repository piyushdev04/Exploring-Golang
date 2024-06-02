package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	piyush := User{"Piyush", "piyush@go.dev", true, 20}
	fmt.Println(piyush)
	fmt.Printf("Piyush details are: %+v\n", piyush)
	piyush.GetStatus()
	piyush.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}