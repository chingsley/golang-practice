package jotters

import (
	"fmt"
	"reflect"
)

func Execute() {
	log := fmt.Println
	log(reflect.TypeOf(log))

	log("Hello, world.")

	var whatToSay string
	var i int
	whatToSay = "Hello again"
	log(whatToSay)

	i = 7
	log("i is set to", i)

	whatWasSaid := saySomething()
	log(whatWasSaid)

	firstName, lastName := getNames()
	log(firstName, lastName)

}

func saySomething() string {
	return "something was said"
}

/**
by the way, go functions can return
more than one thing, wierd!!!!
*/
func getNames() (string, string) {
	return "vin", "Diesel"
}
