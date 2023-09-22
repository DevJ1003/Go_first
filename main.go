package main

import "fmt"

func main() {

	fmt.Println("Hello world!")

	var WhatToSay string
	WhatToSay = "Goodbye, cruel world!"
	fmt.Println(WhatToSay)

	var num int
	num = 65
	fmt.Println("The number is:", num)

	Said, otherSaid := SaySomething()
	fmt.Println("The function SaySomething() says:", Said, otherSaid)

}

func SaySomething() (string, string) {
	return "something", "else"
}
