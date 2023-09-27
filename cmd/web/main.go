package main

import (
	"fmt"
	"net/http"

	"github.com/devj1003/mygoapp/pkg/handlers"
)

// Port number function
const portNumber = ":8081"

// Main func
func main() {

	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
