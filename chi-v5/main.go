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

	corsMw, err := cors.NewMiddleware(cors.Config{
		Origins: []string{"https://example.com"},
		Methods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
		RequestHeaders: []string{"Authorization"},
	})
	if err != nil {
		log.Fatal(err)
	}
	corsMw.SetDebug(true) // turn debug mode on (optional)

	api := chi.NewMux()
	mux.Mount("/api", corsMw.Wrap(api))
	api.Get("/users", handleUsersGet)
	api.Post("/users", handleUsersPost)
	api.Put("/users", handleUsersPut)
	api.Delete("/users", handleUsersDelete)

	log.Fatal(http.ListenAndServe(":8080", mux))
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

func handleUsersPut(w http.ResponseWriter, _ *http.Request) {
	// omitted implementation
}

func handleUsersDelete(w http.ResponseWriter, _ *http.Request) {
	// omitted implementation
}
