package main

import (
	"log"

	mypack "github.com/devj1003/mypackage2/mypack"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := mypack.RandomNumbers(numPool)
	intChan <- randomNumber
}

func main() {

	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
