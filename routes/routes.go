package routes

import (
	"fmt"
	"router-schema/lib"
)

func InitRoutes() *lib.Routes {
	routesList := lib.NewRoutes(nil)
	routesList.Get(lib.Route{
		Handler: func() (any, error) {
			fmt.Println("First handler")
			return nil, nil
		},
		Path: "/alumno/notas/codigo",
	}).Get(lib.Route{
		Handler: func() (any, error) {
			fmt.Println("class room")
			return nil, nil
		},
		Path: "/alumno/class-room",
	})

	return routesList
}
