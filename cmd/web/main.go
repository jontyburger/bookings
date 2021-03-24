package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/jontyburger/bookings/pkg/config"
	"github.com/jontyburger/bookings/pkg/handlers"
	"github.com/jontyburger/bookings/pkg/render"
)

// port number for the server to listen on
const portNumber = ":8080"

// create a variable to store the config.AppConfig
var app config.AppConfig

// init function
func init() {

	// set the if in production or no
	app.InProduction = false
	// set the cache
	app.CacheOn = app.InProduction

	// secure cookie
	b := make([]byte, 32)
	// random read b
	_, err := rand.Read(b)

	// check if err
	if err != nil {
		log.Fatalln(err)
	}

	// make the byte into a base64 string
	str := base64.URLEncoding.EncodeToString(b)

	// create a session and set encryption key
	app.Session = sessions.NewCookieStore([]byte(os.Getenv(str)))

	// set cookie options
	app.Session.Options = &sessions.Options{
		MaxAge:   24 * 60,
		HttpOnly: app.InProduction,
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}
}

// main function
func main() {

	// call the CreateTemplateCache from the render pkg
	tc, err := render.CreateTemplateCache()

	// check if there was an error
	if err != nil {
		log.Fatal("Error creating cache:\n" + err.Error())
	}

	// add the tc to the app.templateCache
	app.TemplateCache = tc

	// create a new repo and return a pointer to the Resporitory
	repo := handlers.NewRepo(&app)

	// Set the Repo
	handlers.SetRepo(repo)

	// // send the app to render.NewTemplate
	render.NewTemplate(&app)

	// setup the server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	// srv listen and serve
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}

}
