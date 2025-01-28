package model

import (
	"github.com/gofiber/fiber/v2"
)

var (
	StatusOK           = fiber.NewError(fiber.StatusOK)
	StatusCreated      = fiber.NewError(fiber.StatusCreated)
	StatusNoContent    = fiber.NewError(fiber.StatusNoContent)
	StatusUnauthorized = fiber.NewError(fiber.StatusUnauthorized)
)

type Response struct {
	Code   int             `json:"code"`
	Status string          `json:"status"`
	Data   any             `json:"data,omitempty"`
	Meta   *MetaPagination `json:"meta,omitempty"`
	Errors any             `json:"errors,omitempty"`
}

type MetaPagination struct {
	Pagination *PaginationMetaData `json:"pagination,omitempty"`
}

type PaginationMetaData struct {
	Total       int             `json:"total"`
	Count       int             `json:"count"`
	PerPage     int             `json:"per_page"`
	CurrentPage int             `json:"current_page"`
	TotalPage   int             `json:"total_page"`
	Links       *PaginationLink `json:"links"`
}

type PaginationLink struct {
	NextPage     string `json:"next_page"`
	PreviousPage string `json:"previous_page"`
}

func WebResponse(ctx *fiber.Ctx, status *fiber.Error, model any) error {
	return ctx.Status(status.Code).JSON(&Response{
		Code:   status.Code,
		Status: status.Message,
		Data:   model,
	})
}

func ErrorResponse(ctx *fiber.Ctx, status *fiber.Error, err any) error {
	return ctx.Status(status.Code).JSON(&Response{
		Code:   status.Code,
		Status: status.Message,
		Errors: err,
	})
}
