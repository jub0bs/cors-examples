package main

import (
	"io"
	"log"
	"net/http"

	"github.com/go-fuego/fuego"
	"github.com/jub0bs/cors"
)

func main() {
	s := fuego.NewServer(
		fuego.WithAddr(":8080"),
	)

	fuego.GetStd(s, "/hello", handleHello) // note: not configured for CORS

	cors, err := cors.NewMiddleware(cors.Config{
		Origins:        []string{"https://example.com"},
		Methods:        []string{http.MethodGet, http.MethodPost},
		RequestHeaders: []string{"Authorization"},
	})
	if err != nil {
		log.Fatal(err)
	}
	cors.SetDebug(true) // turn debug mode on (optional)

	api := fuego.Group(s, "/api")
	fuego.Use(api, cors.Wrap)
	fuego.GetStd(api, "/users", handleUsersGet)
	fuego.PostStd(api, "/users", handleUsersPost)

	s.Run()
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
