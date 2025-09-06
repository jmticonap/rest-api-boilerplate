package routes

import (
	"context"
	"fmt"
	"net/http"
	"router-schema/application/handler"
	"router-schema/lib"
)

// InitRoutes initializes the routes for the application
//
// returns: a pointer to the Routes struct
func InitRoutes() *lib.Routes {
	routesList := lib.NewRoutes(nil)
	routesList.Get(lib.Route{
		Handler: lib.
			NewMiddleware(context.Background()).
			Use(func(w http.ResponseWriter, r *http.Request) (any, error) {
				fmt.Println("Middleware before handler executed")
				return nil, nil
			}).
			After(func(w http.ResponseWriter, r *http.Request) (any, error) {
				fmt.Println("Middleware after handler executed")
				return nil, nil
			}).
			Handler(handler.GreetingIndex),
		Path: "/alumno/notas/codigo",
	}).Get(lib.Route{
		Handler: lib.NewMiddleware(nil).
			Handler(func(w http.ResponseWriter, r *http.Request) (any, error) {
				fmt.Println("Hello from middleware")
				fmt.Println("class room")
				return nil, nil
			}),
		Path: "/alumno/class-room",
	})

	return routesList
}
