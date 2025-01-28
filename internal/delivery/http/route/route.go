package route

import (
	"github.com/dickidarmawansaputra/go-clean-architecture/internal/delivery/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

type RouteConfig struct {
	App            *fiber.App
	Swagger        *swagger.Config
	AuthController *controller.AuthController
}

func Router(config *RouteConfig) {
	route := config.App.Group("/api").Use(recover.New(), cors.New(), healthcheck.New())

	config.UnprotectedRoute(route)
	config.ProtectedRoute(route)
}

func (r *RouteConfig) UnprotectedRoute(route fiber.Router) {
	// swagger route
	route.Get("/docs/*", swagger.New(*r.Swagger))

	// auth routes
	auth := route.Group("/auth")
	auth.Post("/register", r.AuthController.Register)
	auth.Post("/login", r.AuthController.Login)
}

func (r *RouteConfig) ProtectedRoute(route fiber.Router) {

}
