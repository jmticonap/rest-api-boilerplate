package routes

import (
	"fmt"
	"net/http"

	"rest-api/app/application/handler"
	"rest-api/app/lib"
)

// InitRoutes initializes the routes for the application
//
// returns: a pointer to the Routes struct
func InitRoutes() *lib.Routes {
	return lib.NewRoutes().
		Get(lib.Route{
			Handler: lib.NewMiddleware().
				Use(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware before handler executed")
					return nil, nil
				}).
				Use(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware 2 before handler executed")
					return nil, nil
				}).
				After(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware after handler executed")
					return nil, nil
				}).
				Build(handler.GreetingIndex),
			Path: "/alumno/notas/codigo",
		}).
		Get(lib.Route{
			Handler: lib.NewMiddleware().
				Build(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Hello from middleware")
					fmt.Println("class room")
					return nil, nil
				}),
			Path: "/alumno/class-room",
		}).
		Get(lib.Route{
			Handler: nil,
			Path:    "/alumno/class-room/1",
		})
}
