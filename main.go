package main

import (
	"fmt"
	"log"
	"router-schema/lib"
)

func main() {
	log.Default().Println("Creating router")
	router := lib.NewRouter()
	log.Default().Println("Seting up routes")
	router.Get(lib.Route{
		Handler: func() {
			fmt.Println("First handler")
		},
		Path: "/alumno/notas/codigo",
	}).Get(lib.Route{
		Handler: func() {
			fmt.Println("class room")
		},
		Path: "/alumno/class-room",
	})

	log.Default().Println("Getting handler")
	lib.ExeEvent(lib.RequestEvent{
		Method: "GET",
		Path:   "/alumno/notas/codigo",
	}, router.Routes)
	lib.ExeEvent(lib.RequestEvent{
		Method: "GET",
		Path:   "/alumno/class-room",
	}, router.Routes)
}
