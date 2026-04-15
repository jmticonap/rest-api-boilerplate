package handler

import (
	"log"
	"net/http"
	"rest-api/app/lib"
)

func GreetingIndex(r *http.Request) (*lib.MidResponse, error) {
	log.Println("[AFTER] Request processed:", r.Method, r.URL.Path)
	return nil, nil
}
