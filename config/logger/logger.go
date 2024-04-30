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

		// If request body is in form-data, exclude file, extract only key-value pair datas
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
Set logger middleware for REST API, accept both request body & form data.

NOTE : display log data in JSON format
*/
func SetLoggerMiddlewareInJSONFormat(app *fiber.App) {
	var myLogger logger.Config

	switch os.Getenv("MODE") {
	case "prod":
		myLogger = logger.Config{
			Output:     os.Stdout,
			Format:     `{"time": "${time}", "status": "${status}", "method": "${method}", "latency": "${latency}", "ip": "${ip}", "path": "${path}", "query_param": "${queryParams}", "user": "${locals:user}", "request_id": "${locals:requestid}", "request_body": "${customReqBody}", "request_headers": "-", "response_body": "-", "error": "${error}"}` + "\n",
			TimeFormat: "2006/01/02 - 15:04:05",
			// TimeZone:   "Asia/Bangkok",
		}
	default:
		myLogger = logger.Config{
			CustomTags: myCustomLoggerTags,
			Output:     os.Stdout,
			Format:     `{"time": "${time}", "status": "${status}", "method": "${method}", "latency": "${latency}", "ip": "${ip}", "path": "${path}", "query_param": "${queryParams}", "user": "${locals:user}", "request_id": "${locals:requestid}", "request_body": "${customReqBody}", "request_headers": "-", "response_body": "-", "error": "${error}"}` + "\n",
			TimeFormat: "2006/01/02 - 15:04:05",
			// TimeZone:   "Asia/Bangkok",
		}
	}

	(*app).Use(logger.New(myLogger))
}
