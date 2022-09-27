package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/dkr290/web-bookings/internal/config"
	"github.com/dkr290/web-bookings/internal/custerror"
	"github.com/dkr290/web-bookings/internal/handlers"
	"github.com/dkr290/web-bookings/internal/render"
)

const hostPort = "127.0.0.1:8080"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	//changet this to true when going to production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	custerror.FatalErrString("cannot create template cache", err)

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.Newtemplates(&app)

	http.Handle("/favicon.cio", http.NotFoundHandler())
	fmt.Println("starting application on port", hostPort)

	srv := http.Server{
		Addr:    hostPort,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	custerror.FatalErr(err)

}
