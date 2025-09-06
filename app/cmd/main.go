package main

import (
	"log"
	"net/http"
	"router-schema/app/lib"
	"router-schema/app/routes"
)

var routesList *lib.Routes

func init() {
	routesList = routes.InitRoutes()
}

func main() {
	var server http.Server
	server.Addr = "127.0.0.1:3000"
	server.Handler = http.HandlerFunc(
		lib.HttpRouterHandler(routesList.Routes),
	)

	log.Println("Starting HTTP server on", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}
