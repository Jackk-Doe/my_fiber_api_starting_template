package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Jackk-Doe/my_fiber_api_starting_template/api"
	"github.com/Jackk-Doe/my_fiber_api_starting_template/config/dotenv"
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

	// Get a completely configed api applicatioin
	app := api.SetUpAPI()

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
