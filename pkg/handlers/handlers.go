package handlers

import (
	"log"
	"net/http"

	"github.com/jontyburger/bookings/pkg/config"
	"github.com/jontyburger/bookings/pkg/models"
	"github.com/jontyburger/bookings/pkg/render"
)

// create a pointer of the Resporitory struct
var Repo *Repository

// Resporitory struct
type Repository struct {
	App *config.AppConfig
}

// NewRepo functions
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Set Repo function
func SetRepo(r *Repository) {
	Repo = r
}

// Home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	// get the current session
	sess, err := m.App.Session.Get(r, "cookie")

	// check if err
	if err != nil {
		log.Fatalln(err)
	}

	// Set the session
	sess.Values["remote_ip"] = r.RemoteAddr
	// save the session
	if err := sess.Save(r, w); err != nil {
		log.Fatalln(err)
	}

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// get cookie
	sess, err := m.App.Session.Get(r, "cookie")
	// check if error
	if err != nil {
		log.Fatalln(err)
	}

	// set remote_ip string
	rmIP := sess.Values["remote_ip"]

	// make a new StringMap variable of a map array
	var stringMap = make(map[string]string)
	// make a new Data map
	var dataMap = make(map[string]interface{})

	// assign the key and data
	stringMap["Test"] = "Test data for the about page"
	dataMap["remote_ip"] = rmIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		Data:      dataMap,
	})
}
