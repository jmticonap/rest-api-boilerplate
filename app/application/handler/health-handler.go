package handler

import (
	"net/http"
	"rest-api/app/lib"
)

// HealthCheck returns a simple 200 OK status to indicate the service is running.
func HealthCheck(r *http.Request) (*lib.MidResponse, error) {
	return &lib.MidResponse{
		Status: http.StatusOK,
		Data: map[string]string{
			"status": "UP",
		},
	}, nil
}
