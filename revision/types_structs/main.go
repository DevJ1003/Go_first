package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Dev",
		LastName:  "Joshi",
	}

	log.Println(user.FirstName, user.LastName, "BirthDate:", user.BirthDate)
}

// var s = "seven"

// func main() {

// 	var s2 = "six"

// 	log.Println("s is", s)
// 	log.Println("s2 is", s2)

// 	saySomething("test")

// }

// func saySomething(s3 string) (string, string) {

// 	log.Println("s from saySomething func is", s)
// 	return s3, "world"
// }
