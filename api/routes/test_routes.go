package routes

import (
	"github.com/gofiber/fiber/v2"

	testcontrollers "github.com/Jackk-Doe/my_fiber_api_starting_template/api/controllers"
)

func TestRouter(app fiber.Router) {
	app.Get("/", testcontrollers.TestController)
}
