package main

import (
	"log"
	"net/http"
	"time"
	"web3/package/config"
	handlers "web3/package/handlers"

	"github.com/alexedwards/scs/v2"
)

var sessionManager *scs.SessionManager
var app config.AppConfig

func main() {
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = false // need https to set it to true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	app.Session = sessionManager

	repo := handlers.NewRepo(&app)  //pass in reference to our config
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr: ":8080",
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// To run: $ go run ./cmd/web