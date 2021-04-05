package main

import (
	"fmt"
	"net/http"

	"github.com/igolubevJ/golang-project/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Server is runnin on port: %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
