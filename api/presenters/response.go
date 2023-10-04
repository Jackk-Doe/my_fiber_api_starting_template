package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	SUCCESS = 1
	FAIL    = 0
)

func ResponseSuccess(data interface{}) fiber.Map {
	t := time.Now()
	return fiber.Map{
		"timestamp": t.Format("2006-01-02-15-04-05"),
		"status":    SUCCESS,
		"data":      data,
		"error":     nil,
	}
}
