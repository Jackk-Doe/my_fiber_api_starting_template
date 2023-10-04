package routes

import (
	"github.com/gofiber/fiber/v2"

	testcontrollers "<YOUR_GO_MOD_NAME>/api/controllers"
)

func TestRouter(app fiber.Router) {
	app.Get("/", testcontrollers.TestController)
}
