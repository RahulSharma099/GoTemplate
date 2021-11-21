package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RahulSharma099/hello-world/cmd/pkg/config"
	"github.com/RahulSharma099/hello-world/cmd/pkg/handlers"
	"github.com/RahulSharma099/hello-world/cmd/pkg/render"
)

const portNumber = ":8000"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template Cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on PORT: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
