package main

import (
	"context"
	"log"
	"net/http"
	"router-schema/lib"
	"router-schema/routes"
)

var routesList *lib.Routes

func init() {
	routesList = routes.InitRoutes()
}

func main() {
	var server http.Server
	server.Addr = "127.0.0.1:3000"

	server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lib.NewMiddleware(w, r).
			Use(func(ctx context.Context, w http.ResponseWriter, r *http.Request) (any, error) {
				log.Default().Println("[BEFORE] Request received:", r.Method, r.URL.Path)
				return nil, nil
			}).
			After(func(ctx context.Context, w http.ResponseWriter, r *http.Request) (any, error) {
				log.Default().Println("[AFTER] Request processed:", r.Method, r.URL.Path)
				return nil, nil
			}).
			Handler(func(ctx context.Context, w http.ResponseWriter, r *http.Request) (any, error) {
				_, err := lib.HttpRouter(
					lib.RequestEvent{
						Method: r.Method,
						Path:   r.URL.Path,
					},
					routesList.Routes,
				)

				if err != nil {
					log.Printf("Error executing handler: %s\n", err)
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("Some error hapend"))
				} else {
					log.Println("Success execute handler")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("Handler execute successfuly"))
				}

				return nil, nil
			})
	})

	log.Println("Starting HTTP server on", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}
