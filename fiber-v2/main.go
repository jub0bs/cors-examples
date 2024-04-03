package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/jub0bs/cors"
)

func main() {
	app := fiber.New()
	app.Get("/hello", handleHello) // note: not configured for CORS

	// create CORS middleware
	corsMw, err := cors.NewMiddleware(cors.Config{
		Origins:        []string{"https://example.com"},
		Methods:        []string{http.MethodGet, http.MethodPost},
		RequestHeaders: []string{"Authorization"},
	})
	if err != nil {
		log.Fatal(err)
	}
	corsMw.SetDebug(true) // turn debug mode on (optional)

	api := app.Group("/api")
	api.Use(adaptor.HTTPMiddleware(corsMw.Wrap))
	api.Get("/users", handleUsersGet)
	api.Post("/users", handleUsersPost)

	log.Fatal(app.Listen(":8080"))
}

func handleHello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func handleUsersGet(c *fiber.Ctx) error {
	return nil // omitted implementation
}

func handleUsersPost(c *fiber.Ctx) error {
	return nil // omitted implementation
}
