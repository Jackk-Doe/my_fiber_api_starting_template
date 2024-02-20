package logger

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

/*
My Custom logger tags
*/
var myCustomLoggerTags = map[string]logger.LogFunc{

	/// Test custom logger tag
	"myTEST": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
		return output.WriteString("MY_TEST_VALUE")
	},

	/// Custom request body logger tag
	"customReqBody": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
		// If request body is empty, return empty string
		if c.Request().Body() == nil {
			return output.WriteString("")
		}

		// If request body is in form-data
		if contentType := strings.Split(c.Get("Content-Type"), ";")[0]; contentType == "multipart/form-data" {
			form, err := c.MultipartForm()
			if err != nil {
				return output.WriteString("ERROR_GETTING_FORM_DATA")
			}

			var builder strings.Builder
			for key, value := range form.Value {
				if len(value) > 0 && value[0] != "" {
					builder.WriteString(key + "=" + value[0] + "&")
				}
			}
			return output.WriteString(builder.String())
		}

		// If request body is in JSON
		msg := strings.ReplaceAll(string(c.Request().Body()), "\n", "")
		return output.WriteString(msg)
	},
}

/*
Set logger middleware for REST API for JSON data
*/
func SetLoggerMiddlewareJSON(app fiber.Router) {
	var myLogger logger.Config

	/// If in production mode, Display only necessary log datas
	if os.Getenv("MODE") == "prod" {
		myLogger = logger.Config{
			Output: os.Stdout,
			Format: `[API] ${cyan}${time} | ${yellow}${status}${reset} | ${blue}${method}${reset} | ${yellow}${latency}${reset} ` +
				`| IP: ${green}${ip}${reset} | PATH: ${magenta}${path}${reset} | QUERY_PARAM: ${queryParams} | LOCALS: ${cyan}${locals:user} ${reset} | REQUEST_ID: ${locals:requestid} ` +
				`| REQUEST_BODY: ${cyan}- ${reset} | REQUEST_HEADERS: ${cyan}- ${reset} ` +
				`| RESPONSE_BODY: ${cyan}- ${reset} | ERROR: ${red}${error}` + "\n",
			TimeFormat: "2006/01/02 - 15:04:05",
			// TimeZone:   "Asia/Bangkok",
		}
	} else /* /// Else (uat or dev), Display all log datas */
	{
		myLogger = logger.Config{
			CustomTags: myCustomLoggerTags,
			Output:     os.Stdout,
			Format: `[API] ${cyan}${time} | ${yellow}${status}${reset} | ${blue}${method}${reset} | ${yellow}${latency}${reset} ` +
				`| IP: ${green}${ip}${reset} | PATH: ${magenta}${path}${reset} | QUERY_PARAM: ${queryParams} | LOCALS: ${cyan}${locals:user} ${reset} | REQUEST_ID: ${locals:requestid} ` +
				`| REQUEST_BODY: ${cyan}${customReqBody} ${reset} | REQUEST_HEADERS: ${cyan}${header:Authorization} ${reset} ` +
				`| RESPONSE_BODY: ${cyan}${resBody} ${reset} | ERROR: ${red}${error}` + "\n",
			TimeFormat: "2006/01/02 - 15:04:05",
			// TimeZone:   "Asia/Bangkok",
		}
	}

	app.Use(logger.New(myLogger))
}
