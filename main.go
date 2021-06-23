package main

import (
	"github.com/chingsley/testing-go-things/channels"
	"github.com/chingsley/testing-go-things/interfaces"
	"github.com/chingsley/testing-go-things/json_helper"
	"github.com/chingsley/testing-go-things/testing_course"
)

func main() {
	// TESTING INTERFACES
	interfaces.Execute()

	// CREATING A CHANNEL
	channels.Execute()

	// TEST READING AND WRITING JSON
	json_helper.RunMyJsonPackage()

	// RUNNING THE TESTING_COURSE MODULE
	testing_course.Execute()

}
