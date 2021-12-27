package cmd

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ha-dev/getslowestdbqueries/env"
	"github.com/ha-dev/getslowestdbqueries/pkg/middleware"
	"github.com/ha-dev/getslowestdbqueries/pkg/routes"
)

func RunServer() {

	app := fiber.New()

	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.
	routes.PublicRoutes(app)
	routes.NotFoundRoute(app)

	// app.Get("/api/list", Handler.GetQuery)
	log.Fatal(app.Listen(env.ServicePort))

}
