# Parsing HTTP Request Using Echo

This is a simple Go application that demonstrates basic usage of the Echo framework to create a RESTful API for managing user data.

## Getting Started

These instructions will help you set up and run the application on your local machine for development and testing purposes.

### Prerequisites

- Go (v1.16 or higher): [Download and Install Go](https://golang.org/dl/)
- Echo framework (v4): Install it using `go get github.com/labstack/echo/v4`

### Installing

1. Clone the repository:

   ```sh
   git clone https://github.com/umjiiii/parsing-http-request.git
   cd your-go-app
   ```
   
2. Build and run the application.
   
     ```sh
     go run main.go
     ```
The server will start on port 9000. You can access it at `localhost:9000/users`

### Usage

This application exposes a simple RESTful API endpoint for managing user data.

## Create User 
To create a new user, make a POST request to /users with the user's name and email as JSON data:

```shell
curl -X POST http://localhost:9000/users -d '{"first_name": "John", "last_name":"Doe", "email": "john@example.com"}'
```

## Sample Response
The server will respond with the user's data in JSON format:
```json
{
  "first_name": "John",
  "last_name": "Doe",
  "email": "john@example.com"
}
```

## Acknowledgements
This application was developed by following the [Novalagung Parsing HTTP Request Payload with Echo](https://dasarpemrogramangolang.novalagung.com/C-parsing-http-request-payload-echo.html). A big thank you to Novalagung for their tutorial and guidance.

## Contributing
Feel free to contribute to this project by opening issues or pull requests. Any feedback or improvements are highly appreciated.
