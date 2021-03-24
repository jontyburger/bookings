package render

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/jontyburger/bookings/pkg/config"
	"github.com/jontyburger/bookings/pkg/models"
)

var functions template.FuncMap

// make app as a pointer t config.AppConfig
var app *config.AppConfig

// NewTemplate
func NewTemplate(a *config.AppConfig) {
	// set the app as a
	app = a
}

// NewDefaultData function
func NewDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate function
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	// create the type for the tc variable
	tc := map[string]*template.Template{}
	var err error

	// check if the template cache is on
	if app.CacheOn {
		tc = app.TemplateCache
	} else {
		tc, err = CreateTemplateCache()
		// check if has error
		if err != nil {
			log.Fatalln(err)
		}
	}

	// check if the you can get the data from the cache map array
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal(ok)
	}

	// call NewDefaultData
	td = NewDefaultData(td)

	t.Execute(w, td)

}

// CreateTemplateCache function
func CreateTemplateCache() (map[string]*template.Template, error) {

	// create a map array to hold the page name and template
	myCache := map[string]*template.Template{}

	// get all of the pages from the templates folder that end in .page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	// check if an error getting the pages
	if err != nil {
		log.Fatalf("ERROR: You have an error getting the pages:\n" + err.Error())
		return myCache, err
	}

	// create a for loop to go through all of the pages and get the path
	for _, page := range pages {
		// this is taking the file name from the path[page]
		name := filepath.Base(page)

		// create a new template the selected name and parse the files of page
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		// check if an error creating new template
		if err != nil {
			log.Fatalf("ERROR: You have an error creating the new template:\n" + err.Error())
			return myCache, err
		}

		// get the .layout.tmpl files
		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		// check if an error getting layouts
		if err != nil {
			log.Fatalf("ERROR: You have and error trying to get all of the .layout.tmpl pages:\n" + err.Error())
			return myCache, err
		}

		// check if matches is a length more than 0
		if len(matches) > 0 {
			// create a template for the .layout
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

			// check if an error creating the template
			if err != nil {
				log.Fatalf("ERROR: You have an error creating a template for the .layout.tmpl:\n" + err.Error())
			}
		}

		// store the name and ts to the map - myCache
		myCache[name] = ts

	}

	return myCache, nil

}
