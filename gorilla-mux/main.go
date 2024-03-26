package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jub0bs/cors"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", handleHello).
		Methods(http.MethodGet) // note: not configured for CORS

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

	api := router.PathPrefix("/api").Subrouter()
	api.Use(corsMw.Wrap)
	// Note: Because of a design quirk of gorilla/mux,
	// if you add a matcher for some methods,
	// you must (unfortunately) list OPTIONS among them;
	// otherwise, CORS-preflight requests won't reach the CORS middleware,
	// and preflight will invariably fail.
	// See https://github.com/gorilla/mux#handling-cors-requests.
	//
	// For this reason, I recommend using net/http (or Chi, perhaps)
	// instead of gorilla/mux.
	api.HandleFunc("/users", handleUsersGet).
		Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/users", handleUsersPost).
		Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/users", handleUsersPut).
		Methods(http.MethodPut, http.MethodOptions)
	api.HandleFunc("/users", handleUsersDelete).
		Methods(http.MethodDelete, http.MethodOptions)

	log.Fatal(http.ListenAndServe(":8080", router))
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
