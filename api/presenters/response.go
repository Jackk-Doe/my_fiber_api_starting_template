package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	SUCCESS = 1
	FAIL    = 0
)

/*
Return success response
*/
func ResponseSuccess(data interface{}) fiber.Map {
	t := time.Now()
	return fiber.Map{
		"timestamp": t.Format("2006-01-02-15-04-05"),
		"status":    SUCCESS,
		"items":     data,
		"error":     nil,
	}
}

/*
Return list data with pagination infos
*/
// NOTE: if not pagination infos is not required, pass -1 to currentPage, currentPageTotalItem, totalPage
func ResponseSuccessListData(data interface{}, currentPage, currentPageTotalItem, totalPage int) fiber.Map {
	t := time.Now()
	return fiber.Map{
		"timestamp": t.Format("2006-01-02-15-04-05"),
		"status":    SUCCESS,
		"items": fiber.Map{
			"list_data": data,
			"pagination": fiber.Map{
				"current_page":            currentPage,
				"current_page_total_item": currentPageTotalItem,
				"total_page":              totalPage,
			},
		},
		"error": nil,
	}
}
