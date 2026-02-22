package main

import (
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jub0bs/cors"
)

func main() {
	mux := chi.NewMux()
	mux.Get("/hello", handleHello) // note: not configured for CORS

	cors, err := cors.NewMiddleware(cors.Config{
		Origins:        []string{"https://example.com"},
		Methods:        []string{http.MethodGet, http.MethodPost},
		RequestHeaders: []string{"Authorization"},
	})
	if err != nil {
		log.Fatal(err)
	}
	cors.SetDebug(true) // turn debug mode on (optional)

	api := chi.NewMux()
	mux.Mount("/api", cors.Wrap(api))
	api.Get("/users", handleUsersGet)
	api.Post("/users", handleUsersPost)

	if err := http.ListenAndServe(":8080", mux); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func handleHello(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello, World!")
}

func handleUsersGet(w http.ResponseWriter, _ *http.Request) {
	// omitted implementation
}

func handleUsersPost(w http.ResponseWriter, _ *http.Request) {
	// omitted implementation
}
