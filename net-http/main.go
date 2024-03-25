package main

import (
	"io"
	"log"
	"net/http"

	"github.com/jub0bs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", handleHello) // note: not configured for CORS

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

	api := http.NewServeMux()
	mux.Handle("/api/", corsMw.Wrap(api)) // note: method-less pattern here
	api.HandleFunc("GET /api/users", handleUsersGet)
	api.HandleFunc("POST /api/users", handleUsersPost)
	api.HandleFunc("PUT /api/users", handleUsersPut)
	api.HandleFunc("DELETE /api/users", handleUsersDelete)

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
