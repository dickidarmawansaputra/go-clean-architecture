package route

import (
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/delivery/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type RouteConfig struct {
	App            *fiber.App
	Swagger        *swagger.Config
	AuthController *controller.AuthController
}

func Router(config *RouteConfig) {
	route := config.App.Group("/api")

	config.UnprotectedRoute(route)
	config.ProtectedRoute(route)
}

func (r *RouteConfig) UnprotectedRoute(route fiber.Router) {
	// swagger route
	route.Get("/docs/*", swagger.New(*r.Swagger))

	// auth routes
	auth := route.Group("/auth")
	auth.Post("/register", r.AuthController.Register)
}

func (r *RouteConfig) ProtectedRoute(route fiber.Router) {

}
