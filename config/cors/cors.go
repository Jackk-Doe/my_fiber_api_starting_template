package cors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetCORSMiddleware(app *fiber.App) {
	app.Use(cors.New(cors.Config{}))
}
