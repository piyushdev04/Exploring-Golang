package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// var declared without a coresponding initialization are zero-valued 0
	var e int
	fmt.Println(e)

	// ':=' syntax is shorthand for declaring and initializing a variable
	f := "BATMAN"
	fmt.Println(f)
}
