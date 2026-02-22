package main

import (
	"log"
	"net/http"

	"github.com/jub0bs/cors"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/hello", handleHello) // note: not configured for CORS

	// create CORS middleware
	cors, err := cors.NewMiddleware(cors.Config{
		Origins:        []string{"https://example.com"},
		Methods:        []string{http.MethodGet, http.MethodPost},
		RequestHeaders: []string{"Authorization"},
	})
	if err != nil {
		log.Fatal(err)
	}
	cors.SetDebug(true) // turn debug mode on (optional)

	api := e.Group("/api", echo.WrapMiddleware(cors.Wrap))
	api.GET("/users", handleUsersGet)
	api.POST("/users", handleUsersPost)

	log.Fatal(e.Start(":8080"))
}

func handleHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func handleUsersGet(c echo.Context) error {
	return nil // omitted implementation
}

func handleUsersPost(c echo.Context) error {
	return nil // omitted implementation
}
