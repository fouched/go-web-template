package main

import (
	"github.com/fouched/go-web-template/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Instance.Home)

	//mux.Route("/caregivers", func(r chi.Router) {
	//	r.Get("/", handlers.Instance.ContactsListGet)
	//	r.Post("/", handlers.Instance.ContactsListPost)
	//	r.Get("/new", handlers.Instance.ContactsNewGet)
	//	r.Post("/new", handlers.Instance.ContactsNewPost)
	//	r.Get("/{id}", handlers.Instance.ContactsViewGet)
	//	r.Get("/{id}/edit", handlers.Instance.ContactsEditGet)
	//	r.Post("/{id}/edit", handlers.Instance.ContactsEditPost)
	//	r.Post("/delete", handlers.Instance.ContactsDeleteSelected)
	//	r.Delete("/{id}", handlers.Instance.ContactsDelete)
	//	r.Get("/{id}/email", handlers.Instance.ContactsEmailValidation)
	//	r.Get("/archive", handlers.Instance.ArchiveGet)
	//	r.Post("/archive", handlers.Instance.ArchivePost)
	//})

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
