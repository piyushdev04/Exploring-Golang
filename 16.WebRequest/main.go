package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://bento.me/piyush-dev"

func main() {
	fmt.Println("Go web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
