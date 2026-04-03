package routes

import (
	"context"
	"fmt"
	"net/http"

	"rest-api/app/application/handler"
	"rest-api/app/lib"
)

// InitRoutes initializes the routes for the application
//
// returns: a pointer to the Routes struct
func InitRoutes() *lib.Routes {
	ctx := context.Background()
	return lib.NewRoutes().Get(lib.Route{
		Handler: lib.
			NewMiddleware(ctx).
			Use(func(w http.ResponseWriter, r *http.Request) (any, error) {
				fmt.Println("Middleware before handler executed")
				return nil, nil
			}).
			Use(func(w http.ResponseWriter, r *http.Request) (any, error) {
				fmt.Println("Middleware 2 before handler executed")
				return nil, nil
			}).
			After(func(w http.ResponseWriter, r *http.Request) (any, error) {
				fmt.Println("Middleware after handler executed")
				return nil, nil
			}).
			Handler(handler.GreetingIndex),
		Path: "/alumno/notas/codigo",
	}).Get(lib.Route{
		Handler: lib.NewMiddleware(ctx).
			Handler(func(w http.ResponseWriter, r *http.Request) (any, error) {
				fmt.Println("Hello from middleware")
				fmt.Println("class room")
				return nil, nil
			}),
		Path: "/alumno/class-room",
	}).Get(lib.Route{
		Handler: nil,
		Path:    "/alumno/class-room/1",
	})
}
