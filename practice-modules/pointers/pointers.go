package pointers

import "log"

func Execute() {
	var myString string

	myString = "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after func call, myString is set to", myString)
}

func changeUsingPointer(s *string) {
	log.Println(s)
	newValue := "Red"

	*s = newValue
}
