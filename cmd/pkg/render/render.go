package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/RahulSharma099/hello-world/cmd/pkg/config"
	"github.com/RahulSharma099/hello-world/cmd/pkg/handlers"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *handlers.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		//get the templateCache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, erro := buf.WriteTo(w)
	if erro != nil {
		fmt.Println("Error writing template to a browser", erro)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		Ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			Ts, err := Ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
			myCache[name] = Ts
		}

	}
	return myCache, nil
}
