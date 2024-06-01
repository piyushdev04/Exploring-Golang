package main

import "fmt"

func main() {

	// String
	fmt.Println("go" + "lang")

	// Integer and float
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	// Booleans
	// '&&' AND opt.
	fmt.Println(true && false)
	fmt.Println(false && false)
	fmt.Println(true && true)

	// '||' OR opt.
	fmt.Println(true || false)
	fmt.Println(false || false)
	fmt.Println(true || true)

	// '!'
	fmt.Println(!true)
	fmt.Println(!false)
}
