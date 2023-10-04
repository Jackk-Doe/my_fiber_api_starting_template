package controllers

import (
	"github.com/gofiber/fiber/v2"

	"<YOUR_GO_MOD_NAME>/api/presenters"
)

func TestController(c *fiber.Ctx) error {
	test := "Hello World"
	res := presenters.ResponseSuccess(test)
	return c.Status(fiber.StatusOK).JSON(res)
}
