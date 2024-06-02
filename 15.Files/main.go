package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in file - I'm Batman!"

	file, err := os.Create("./mylcofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	lenght, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is: ", lenght)
	defer file.Close()
	readFile("./mylcofile.txt")
}

func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}