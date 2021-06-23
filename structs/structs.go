package structs

import (
	"fmt"
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

func (user *User) getFullName() string {
	return user.FirstName + " " + user.LastName
}

func Execute() {
	user1 := User{
		FirstName:   "Vin",
		LastName:    "Diesel",
		PhoneNumber: "08208208 20802803",
	}

	var user2 User
	user2.FirstName = "Mare"
	user2.LastName = "Shoiban"

	log.Println(user1)
	log.Println("\nFirstName: ", user1.FirstName, "\nLastName: ", user1.LastName, "\nPhoneNumber: ", user1.PhoneNumber, "\nAge: ", user1.Age, "\nBirthDate: ", user1.BirthDate)

	fmt.Println(user1.getFullName())
	fmt.Println(user2.getFullName())
}
