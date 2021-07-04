package main

import (
	"encoding/gob"
	"fmt"
	config2 "github.com/adamszpilewicz/bookings/internal/config"
	handlers2 "github.com/adamszpilewicz/bookings/internal/handlers"
	"github.com/adamszpilewicz/bookings/internal/models"
	render2 "github.com/adamszpilewicz/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config2.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {

	// what I am going to put into session
	gob.Register(models.Reservation{})

	app.InProduction = false
	session = scs.New()
	session.Lifetime = time.Minute * 30
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render2.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers2.NewRepo(&app)
	handlers2.NewHandlers(repo)

	render2.NewTemplates(&app)

	return nil
}
