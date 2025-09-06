package handler

import (
	"log"
	"net/http"
)

func GreetingIndex(w http.ResponseWriter, r *http.Request) (any, error) {
	log.Println("[AFTER] Request processed:", r.Method, r.URL.Path)
	return nil, nil
}
