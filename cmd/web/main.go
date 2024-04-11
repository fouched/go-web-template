package main

import (
	"database/sql"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/fouched/go-web-template/internal/config"
	"github.com/fouched/go-web-template/internal/handlers"
	"github.com/fouched/go-web-template/internal/render"
	"github.com/fouched/go-web-template/internal/repo"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"net/http"
	"time"
)

const port = ":9080"
const dbString = "host=localhost port=5432 dbname=go-web-template user=go password=password"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	dbPool, err := run()
	if err != nil {
		log.Fatalln(err)
	}

	//seed(dbPool)

	// we have database connectivity, close it after app stops
	defer dbPool.Close()

	srv := &http.Server{
		Addr:    port,
		Handler: routes(),
	}
	fmt.Println(fmt.Sprintf("Starting application on %s", port))

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func run() (*sql.DB, error) {
	dbPool, err := repo.CreateDbPool(dbString)
	if err != nil {
		log.Fatal("Cannot connect to database! Dying argh...")
	}

	// register complex type for session
	// ... nothing yet

	// create the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	app.InProduction = false

	hc := handlers.NewConfig(&app)
	handlers.NewHandlers(hc)
	render.NewRenderer(&app)

	return dbPool, nil
}
