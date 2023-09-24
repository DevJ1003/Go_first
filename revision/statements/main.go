package main

import "log"

func main() {

	myVar := "horse"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")
	default:
		log.Println("cat is something else!")
	}
}

// func main() {

// 	myNum := 100
// 	boolean_dt := false

// 	if myNum > 99 && !boolean_dt {
// 		log.Println("myNum is greater than 99 and boolean is set to true!")
// 	} else if myNum < 100 && boolean_dt {
// 		log.Println("1")
// 	} else if myNum == 101 || boolean_dt {
// 		log.Println("2")
// 	} else if myNum > 1000 && boolean_dt == false {
// 		log.Println("3")
// 	}
// }
