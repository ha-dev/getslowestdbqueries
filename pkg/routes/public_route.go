package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/ha-dev/getslowestdbqueries/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Post("List", controllers.GetQuerys)                            // get list of all slowest queries
	route.Post("AddQueryForCheck", controllers.AddQuerysForCheckSlowest) // Get Querys For Check Slowest

}
