package main

import "github.com/chingsley/testing-go-things/server1"

// "github.com/chingsley/testing-go-things/channels"
// "github.com/chingsley/testing-go-things/http_module"
// "github.com/chingsley/testing-go-things/interfaces"
// "github.com/chingsley/testing-go-things/json_helper"
// "github.com/chingsley/testing-go-things/testing_course"

func main() {
	// // TESTING INTERFACES
	// interfaces.Execute()

	// // CREATING A CHANNEL
	// channels.Execute()

	// // TEST READING AND WRITING JSON
	// json_helper.RunMyJsonPackage()

	// // RUNNING THE TESTING_COURSE MODULE
	// testing_course.Execute()

	// // RUNNING MY HTTP MODULE
	// http_module.Execute()

	server1.MyServerRun()

}

// import (
// 	"fmt"
// 	"html/template"
// 	"net/http"
// )

// const portNumber = ":8080"

// func Home(w http.ResponseWriter, r *http.Request) {
// 	renderTemplate(w, "home.page.tmpl")
// }

// func About(w http.ResponseWriter, r *http.Request) {
// }

// func main() {
// 	http.HandleFunc("/", Home)
// 	http.HandleFunc("/about", About)

// 	fmt.Println("Starting the applicatin on port: ", portNumber)
// 	_ = http.ListenAndServe(portNumber, nil)
// }

// func renderTemplate(w http.ResponseWriter, tmpl string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("error parsing template: ", err)
// 		return
// 	}
// }
