package api

import (
	"errors"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"

	testRouter "github.com/Jackk-Doe/my_fiber_api_starting_template/api/routes"
	"github.com/Jackk-Doe/my_fiber_api_starting_template/api/validators"
	"github.com/Jackk-Doe/my_fiber_api_starting_template/config/cors"
	"github.com/Jackk-Doe/my_fiber_api_starting_template/config/logger"
	requestid "github.com/Jackk-Doe/my_fiber_api_starting_template/config/requestId"
)

func SetUpAPI() *fiber.App {

	var apiName = os.Getenv("API_NAME")
	var apiVersion = os.Getenv("API_VERSION")
	var mode = os.Getenv("MODE")
	var buildAt = os.Getenv("BUILD_DATE")
	var startRunAt = time.Now().Format("2006-01-02 15:04:05")

	/// My custom config for Fiber
	myConfig := fiber.Config{
		// DisableStartupMessage: true,
		AppName: apiName,
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			// Send custom error response
			err = ctx.Status(code).JSON(fiber.Map{
				"timestamp": time.Now().Format("2006-01-02-15-04-05"),
				"status":    0,
				"items":     nil,
				"error":     err.Error(),
			})

			// Return from handler
			return err
		},
	}

	app := fiber.New(myConfig)

	// Set up middlewares & configs
	cors.SetCORSMiddleware(app)
	requestid.SetRequestIdMiddleware(app)
	logger.SetLoggerMiddlewareInJSONFormat(app)

	// Set up other third-party packages (firebase, db, graphql ... etc)
	validators.Init()

	/// API Routes
	api := app.Group(os.Getenv("API_PREFIX"))

	// API Info path
	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"API_NAME":     apiName,
			"API_VERSION":  apiVersion,
			"MODE":         mode,
			"BUILD_AT":     buildAt,
			"START_RUN_AT": startRunAt,
		})
	})

	// healthz route
	api.Get("/healthz", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "OK",
		})
	})

	/**
	* API Routes
	**/
	// TODO: Add your API routes here

	// Test route
	testApi := api.Group("/test")
	testRouter.TestRouter(testApi)

	// Return a complete api app
	return app
}
