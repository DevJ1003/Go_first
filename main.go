package main

import (
	"errors"
	"fmt"
	"net/http"
)

// Port number function
const portNumber = ":8081"

// HomePage func for home page handler
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! Home Page here!")
}

// AboutPage func for about page handler
func AboutPage(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is about page and 2 + 2 is: %d", sum))
	// fmt.Fprintf(w, "About Page!")
}

// addValues func adds two integers and returns the sum
func addValues(x, y int) int {
	return x + y
}

// DividePage func for divide page handler
func DividePage(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(10.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 10.0, 0.0, f))
}

// divideValues func to divide two values
func divideValues(a, b float32) (float32, error) {

	if b <= 0 {
		err := errors.New("Cannot divide by zero!")
		return 0, err
	}

	result := a / b
	return result, nil
}

// Main func
func main() {

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)
	http.HandleFunc("/divide", DividePage)

	fmt.Println(fmt.Sprintf("Strting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
