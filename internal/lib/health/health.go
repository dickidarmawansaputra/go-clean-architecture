package health

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type FiberStatus struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Info   Info   `json:"info,omitempty"`
}

type Info struct {
	Title   string `json:"title"`
	Version string `json:"version"`
	Docs    string `json:"docs"`
}

func Status(ctx *fiber.Ctx) error {
	swaggerAgent := fiber.Get(fmt.Sprintf("%s/api/docs/doc.json", ctx.BaseURL()))
	_, infoBody, _ := swaggerAgent.Bytes()

	var swagger *FiberStatus
	_ = json.Unmarshal(infoBody, &swagger)

	info := Info{
		Title:   swagger.Info.Title,
		Version: swagger.Info.Version,
		Docs:    fmt.Sprintf("%s/api/docs", ctx.BaseURL()),
	}

	agent := fiber.Get(fmt.Sprintf("%s/api/livez", ctx.BaseURL()))
	statusCode, _, errs := agent.Bytes()
	if len(errs) > 0 {
		return ctx.Status(statusCode).JSON(&FiberStatus{
			Code:   fiber.NewError(statusCode).Code,
			Status: fiber.NewError(statusCode).Message,
			Info:   info,
		})
	}

	return ctx.Status(statusCode).JSON(&FiberStatus{
		Code:   fiber.NewError(statusCode).Code,
		Status: fiber.NewError(statusCode).Message,
		Info:   info,
	})
}
