package main

import (
	"basic-golang-crud/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, world!")
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
