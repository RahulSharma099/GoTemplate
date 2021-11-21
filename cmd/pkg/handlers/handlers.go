package handlers

import (
	"net/http"

	"github.com/RahulSharma099/hello-world/cmd/pkg/config"
	"github.com/RahulSharma099/hello-world/cmd/pkg/models"
	"github.com/RahulSharma099/hello-world/cmd/pkg/render"
)

//Repo is the repository used by handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

//About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic\
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!!"

	//send some data to the template

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
