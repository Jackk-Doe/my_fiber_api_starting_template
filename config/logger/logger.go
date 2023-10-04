package logger

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetLoggerMiddleware(app *fiber.App) {

	/// OPTION 1 - Display all logs
	// My preferrenced Logger config
	// myLogger := logger.Config{
	// 	Output: os.Stdout,
	// 	Format: `[API] ${cyan}${time} |${yellow}${status}${reset}|${blue}${method}${reset} | ${yellow}${latency}${reset} ` +
	// 		`| IP: ${green}${ip}${reset} | PATH: ${magenta}${path}${reset} | REQUEST_ID: ${locals:requestid} | QUERY_PARAM: ${queryParams} ` +
	// 		`| REQUEST_BODY: ${cyan}${body} ${reset} | REQUEST_HEADERS: ${cyan}${header:Authorization} ${reset} ` +
	// 		`| RESPONSE_BODY: ${cyan}${resBody} ${reset} | ERROR ${red}${error}` + "\n",
	// 	TimeFormat: "2006/01/02 - 15:04:05",
	// 	// TimeZone:   "Asia/Bangkok",
	// }

	/// OPTION 2 - Display only necessary logs (not include request body and response body)
	myLogger := logger.Config{
		Output: os.Stdout,
		Format: `[API] ${cyan}${time} | ${yellow}${status}${reset} | ${blue}${method}${reset} | ${yellow}${latency}${reset} ` +
			`| IP: ${green}${ip}${reset} | PATH: ${magenta}${path}${reset} | QUERY_PARAM: ${queryParams} | REQUEST_ID: ${locals:requestid} ` +
			`| REQUEST_BODY: ${cyan}- ${reset} | REQUEST_HEADERS: ${cyan}${header:Authorization} ${reset} ` +
			`| RESPONSE_BODY: ${cyan}- ${reset} | ERROR ${red}${error}` + "\n",
		TimeFormat: "2006/01/02 - 15:04:05",
		// TimeZone:   "Asia/Bangkok",
	}

	app.Use(logger.New(myLogger))
}
