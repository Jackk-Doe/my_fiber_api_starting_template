# my_fiber_api_starting_template
A template for Go Fiber API project by Jackk-Doe.<br>

- Create basic info route.<br>
- Set up basic CORS, logger, request-id assigning.<br> 
- Create a folder structures for the API.<br>
- Create a custom success response JSON.<br>
- Override custom fail response JSON.<br>
- Separate Development and Production mode. (DEV, UAT, BUILD, TEST & PROD)<br>


## Initialise go mod and install necessary go packages (fiber & dotenv) 
```
go mod init <YOUR_GO_MOD_NAME>
```

Install packages :
```
go get github.com/gofiber/fiber/v2
go get github.com/joho/godotenv
go get github.com/go-playground/validator/v10
```

Or just with : `go mod tidy`

## Create .env file, use .env.example file as a reference.

## Replace <YOUR_GO_MOD_NAME> with a go module project name in files : `main.go`<br>


## Run command
`$ go run cmd/api/main.go`

Or debug via CompileDaemon with: <br>
`$ CompileDaemon -build="go build -o ./build/api cmd/api/main.go" -command="./build/api"`


## Test routes 
```
localhost:3000/api/v1

OR

localhost:3000/api/v1/test
```


## Success & Fail Response format
SUCCESS Response body
```
{
    "items": {
        "something": "something"
    },
    "error": null,
    "status": 1,
    "timestamp": "2023-10-05-11-21-41"
}
```

FAIL Response body
```
{
    "items": null,
    "error": "<ERROR_MESSAGE>",
    "status": 0,
    "timestamp": "2023-10-05-14-55-03"
}
```

