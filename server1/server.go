package server1

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
}

func MyServerRun() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Starting the applicatin on port: ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}
