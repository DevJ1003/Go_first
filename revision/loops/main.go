package main

import "log"

func main() {

	// ==================================================================
	// animals := []string{"dog", "fish", "horse", "cow", "cat"}
	// for i, animal := range animals {
	// 	log.Println(i, animal)
	// }

	// ==================================================================
	// animals := []string{"dog", "fish", "horse", "cow", "cat"}
	// for _, animal := range animals {
	// 	log.Println(animal)
	// }

	// ==================================================================
	// animals := make(map[string]string)
	// animals["dog"] = "Fido"
	// animals["fish"] = "Golden"

	// // for _, animal := range animals {
	// // 	log.Println(animal)
	// // }
	// for animalType, animal := range animals {
	// 	log.Println(animalType, animal)
	// }

	// ==================================================================
	// var FirstLine = "Once upon a time"

	// for i, l := range FirstLine {
	// 	log.Println(i, ":", l)
	// }

	// ==================================================================
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Silva", "john@silva.com", 30})
	users = append(users, User{"Mike", "Shindoa", "mike@shinoda.com", 20})
	users = append(users, User{"Chris", "Morris", "chris@morris.com", 40})
	users = append(users, User{"Maruti", "Suzuki", "maruti@suzuki.com", 18})

	for i, l := range users {
		log.Println(i, l.FirstName, l.LastName, l.Email, l.Age)
	}
	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
