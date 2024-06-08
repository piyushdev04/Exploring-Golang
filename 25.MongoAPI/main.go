package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/piyushdev04/mongoapi/router"
)

func main() {
	fmt.Println("Mongo API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("listening at port 4000 ...")
}
