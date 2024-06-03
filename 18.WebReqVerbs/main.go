package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content lenght is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	
	//fmt.Println(string(content))
}
