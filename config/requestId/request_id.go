package requestid

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func SetRequestIdMiddleware(app *fiber.App) {
	app.Use(requestid.New())
}
