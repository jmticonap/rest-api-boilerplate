# Go REST API Boilerplate

A simple and lightweight REST API boilerplate written in Go, featuring a custom router implementation. This project provides a basic structure for building RESTful services with a focus on a clear and organized routing mechanism.

## Features

-   Custom router implementation using a tree-like structure.
-   Supports standard HTTP methods: `GET`, `POST`, `PUT`, `DELETE`, `PATCH`, `OPTIONS`, `HEAD`, `CONNECT`, `TRACE`.
-   Chainable route definitions.
-   Basic event handling for incoming requests.

## Getting Started

### Prerequisites

-   Go 1.21 or higher

### Installation

1.  Clone the repository:
    ```sh
    git clone https://github.com/jmticonap/rest-api-boilerplate
    cd rest-api-boilerplate
    ```

2.  Install dependencies:
    ```sh
    go mod tidy
    ```

### Running the Application

To run the main application, which demonstrates a few sample routes:

```sh
go run main.go
```

You should see output indicating that the router is being created and the handlers for the defined routes are being executed.

## Usage

Here's a basic example of how to define routes using the custom router:

```go
package main

import (
	"fmt"
	"router-schema/lib"
)

func main() {
	// Create a new router
	router := lib.NewRouter()

	// Define a GET route
	router.Get(lib.Route{
		Handler: func() {
			fmt.Println("Handler for /users")
		},
		Path: "/users",
	})

    // Define a nested GET route
	router.Get(lib.Route{
		Handler: func() {
			fmt.Println("Handler for /users/profile")
		},
		Path: "/users/profile",
	})

	// Execute an event to trigger a route handler
	lib.ExeEvent(lib.RequestEvent{
		Method: "GET",
		Path:   "/users/profile",
	}, router.Routes)
}
```

## Testing

This project uses the standard `go test` command and includes a `Makefile` for convenience. To run the tests and generate a coverage report:

```sh
make test
```

This command will:
1.  Run all tests in verbose mode.
2.  Generate a `coverage.out` file with the test coverage profile.
3.  Open the HTML coverage report in your default browser.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
