package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes, myMessage := proAdder(2,3,4,5)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro message is: ", myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func greeter() {
	fmt.Println("Namstey from golang")
}

func proAdder(value ...int) (int, string){
	total := 0

	for _, val := range value{
		total += val 
	}

	return total, "Hi Pro result function"

}