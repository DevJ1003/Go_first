package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devj1003/mygoapp/pkg/config"
	"github.com/devj1003/mygoapp/pkg/handlers"
	"github.com/devj1003/mygoapp/pkg/render"
)

// Port number function
const portNumber = ":8081"

// Main func
func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		// fmt.Println(err)
		log.Fatal("Cannot create template!")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.HomePage)
	http.HandleFunc("/about", handlers.Repo.AboutPage)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
