package lib_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rest-api/app/lib"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpRouterHandler(t *testing.T) {
	ctx := context.Background()
	routes := lib.NewRoutes().Get(lib.Route{
		Handler: lib.
			NewMiddleware(ctx).
			Handler(func(w http.ResponseWriter, r *http.Request) (any, error) {
				result := map[string]string{
					"message": "Hello, World!",
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				if err := json.NewEncoder(w).Encode(result); err != nil {
					return nil, err
				}

				return result, nil
			}),
		Path: "/some/path",
	})

	req := httptest.NewRequest("GET", "http://example.com/some/path", nil)
	rec := httptest.NewRecorder()
	lib.HttpRouterHandler(routes.Routes)(rec, req)

	assert.EqualValues(t, http.StatusOK, rec.Code)
	assert.EqualValues(t, "application/json", rec.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.EqualValues(t, "Hello, World!", response["message"])
}

func TestHttpRouterHandlerNotFound(t *testing.T) {
	ctx := context.Background()
	routes := lib.NewRoutes().Get(lib.Route{
		Handler: lib.
			NewMiddleware(ctx).
			Handler(func(w http.ResponseWriter, r *http.Request) (any, error) {
				result := map[string]string{
					"message": "Hello, World!",
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				if err := json.NewEncoder(w).Encode(result); err != nil {
					return nil, err
				}

				return result, nil
			}),
		Path: "/some/path",
	})

	req := httptest.NewRequest("GET", "http://example.com/wrong/path", nil)
	rec := httptest.NewRecorder()
	lib.HttpRouterHandler(routes.Routes)(rec, req)

	assert.EqualValues(t, http.StatusNotFound, rec.Code)
	assert.EqualValues(t, "application/json", rec.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.EqualValues(t, "Path not found", response["message"])
}

func TestHttpRouterHandlerNil(t *testing.T) {
	routes := lib.NewRoutes().Get(lib.Route{
		Handler: nil,
		Path:    "/some/path",
	})

	req := httptest.NewRequest("GET", "http://example.com/some/path", nil)
	rec := httptest.NewRecorder()
	lib.HttpRouterHandler(routes.Routes)(rec, req)

	assert.EqualValues(t, http.StatusNotFound, rec.Code)
	assert.EqualValues(t, "application/json", rec.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.EqualValues(t, "Path not found", response["message"])
}

func TestHttpRouterHandlerExecBefore(t *testing.T) {
	ctx := context.Background()
	routes := lib.NewRoutes().Get(lib.Route{
		Handler: lib.
			NewMiddleware(ctx).
			Handler(func(w http.ResponseWriter, r *http.Request) (any, error) {
				result := map[string]string{
					"message": "Hello, World!",
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				if err := json.NewEncoder(w).Encode(result); err != nil {
					return nil, err
				}

				return result, nil
			}),
		Path: "/some/path",
	})

	req := httptest.NewRequest("GET", "http://example.com/wrong/path", nil)
	rec := httptest.NewRecorder()
	lib.HttpRouterHandler(routes.Routes)(rec, req)

	assert.EqualValues(t, http.StatusNotFound, rec.Code)
	assert.EqualValues(t, "application/json", rec.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.EqualValues(t, "Path not found", response["message"])
}
