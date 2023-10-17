# my_fiber_api_starting_template
A template for Go Fiber API project by Jackk-Doe.<br>

- Create basic info route.<br>
- Set up basic CORS, logger, request-id assigning.<br> 
- Create a folder structures for the API.<br>
- Create a custom success response JSON.<br>
- Override custom fail response JSON.<br>
- Separate Development and Production mode. (DEV & PROD)<br>

## Initialise go mod and install necessary go packages (fiber & dotenv) 
```
go mod init <YOUR_GO_MOD_NAME>
go get github.com/gofiber/fiber/v2
go get github.com/joho/godotenv
```
NOTE: .env file must be created.<br>
NOTE: also replace <YOUR_GO_MOD_NAME> with a go module project name in files : `main.go`<br>

## Run command
`go run cmd/api/main.go`

Or debug via CompileDaemon with: <br>
`CompileDaemon -build="go build -o ./build/api cmd/api/main.go" -command="./build/api"`