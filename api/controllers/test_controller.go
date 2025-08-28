package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Jackk-Doe/my_fiber_api_starting_template/api/presenters"
)

func TestController(c *fiber.Ctx) error {
	test := "Hello World"
	res := presenters.ResponseSuccess(test)
	return c.Status(fiber.StatusOK).JSON(res)
}
