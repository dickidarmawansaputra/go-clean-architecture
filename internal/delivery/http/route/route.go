package route

import (
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/delivery/controller"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	AuthController *controller.AuthController
}

func Router(config *RouteConfig) {
	route := config.App.Group("/api")

	config.UnprotectedRoute(route)
	config.ProtectedRoute(route)
}

func (r *RouteConfig) UnprotectedRoute(route fiber.Router) {
	auth := route.Group("/auth")

	auth.Post("/register", r.AuthController.Register)
}

func (r *RouteConfig) ProtectedRoute(route fiber.Router) {

}
