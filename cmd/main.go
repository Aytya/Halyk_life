package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "httproxy/docs"
	"httproxy/internal/handler"
	"log"
	"net/http"
)

// @title          Proxy Server API
// @version         1.22.4
// @description    This server handles incoming HTTP requests from clients, forwards them to external services, and returns the result to the client in JSON format.

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Post("/proxy", handler.HandleRequest)
	r.Get("/stored", handler.GetStoredRequestHandler)

	log.Println("Starting HTTP server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
