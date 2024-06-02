package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	piyush := User{"Piyush", "piyush@go.dev", true, 20}
	fmt.Println(piyush)
	fmt.Printf("Piyush details are: %+v\n", piyush)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
