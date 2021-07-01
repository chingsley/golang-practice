package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chingsley/basic-web-app-with-go/pkg/config"
	"github.com/chingsley/basic-web-app-with-go/pkg/handlers"
	"github.com/chingsley/basic-web-app-with-go/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting the applicatin on port: ", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}

/*
USE THE COMMAND BELOW TO RUN THE APP
 go run cmd/web/*.go
*/
