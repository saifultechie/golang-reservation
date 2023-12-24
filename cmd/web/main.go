package main

import (
	"encoding/gob"
	"fmt"
	"github/saifultechie/booking/pkg/config"
	"github/saifultechie/booking/pkg/handlers"
	"github/saifultechie/booking/pkg/renders"
	models "github/saifultechie/booking/pkg/templateData"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
	routes := Routes(&app)
	http.ListenAndServe(":8181", routes)
	fmt.Println("serve is running")

}

func run() error {
	// what kind of data put in session
	gob.Register(models.Reservation{})

	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := renders.CreateTemplateCache()

	if err != nil {
		log.Fatal("can not create template cache")
		return err
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	renders.NewTemplate(&app)

	return nil
}
