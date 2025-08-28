package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"

	testRouter "github.com/Jackk-Doe/my_fiber_api_starting_template/api/routes"
	"github.com/Jackk-Doe/my_fiber_api_starting_template/api/validators"
	"github.com/Jackk-Doe/my_fiber_api_starting_template/config/cors"
	"github.com/Jackk-Doe/my_fiber_api_starting_template/config/dotenv"
	"github.com/Jackk-Doe/my_fiber_api_starting_template/config/logger"
	requestid "github.com/Jackk-Doe/my_fiber_api_starting_template/config/requestId"
)

func init() {
	// NOTE: MODE has 'dev', 'uat', 'build', 'test', 'prod' or '' (empty) values
	mode := os.Getenv("MODE")

	if mode == "" {
		// Set up environment variables from .env file
		dotenv.SetDotenv()
	}

	// In prod or uat mode, load environment variables from system
	log.Println("------ Running in '" + mode + "' mode... ------")
}

func main() {
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

	// Set up middleware
	cors.SetCORSMiddleware(app)
	requestid.SetRequestIdMiddleware(app)

	// Set up other third-party packages
	validators.Init()

	/// API Routes
	api := app.Group(os.Getenv("API_PREFIX"))

	/// NOTE: Define Info & Healthz routes before setting up logger middleware, so that they are not logged
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

	// Set up logger middleware
	logger.SetLoggerMiddlewareInJSONFormat(api)

	/// API Routes
	// TODO: Add your API routes here

	// Test route
	testApi := api.Group("/test")
	testRouter.TestRouter(testApi)

	// Run server in a separate goroutine so it doesn't block
	go func() {
		if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
			log.Panic(err)
		}
	}()

	// Create channel to signify a signal being sent
	c := make(chan os.Signal, 1)

	// When an interrupt or termination signal is sent, notify the channel
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c // This blocks the main thread until an interrupt is received
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	log.Println("Running cleanup tasks...")
	// Your cleanup tasks go here ...

	log.Println("Fiber was successful shutdown.")
}
