package route

import (
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App *fiber.App
}

func Router(config *RouteConfig) {
	route := config.App.Group("/api")
	UnprotectedRoute(config, route)
	ProtectedRoute(config, route)
}

func UnprotectedRoute(config *RouteConfig, route fiber.Router) {

}

func ProtectedRoute(config *RouteConfig, route fiber.Router) {

}
