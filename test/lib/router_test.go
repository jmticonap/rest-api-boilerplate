package lib_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rest-api/app/lib"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpRouterHandler(t *testing.T) {
	routes := lib.NewRoutes().
		Get(lib.Route{
			Handler: lib.NewMiddleware().
				Build(func(r *http.Request) (*lib.MidResponse, error) {
					result := map[string]string{
						"message": "Hello, World!",
					}

					return &lib.MidResponse{
						Status: http.StatusOK,
						Data:   result,
					}, nil
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
	routes := lib.NewRoutes().Get(lib.Route{
		Handler: lib.
			NewMiddleware().
			Build(func(r *http.Request) (*lib.MidResponse, error) {
				result := map[string]string{
					"message": "Hello, World!",
				}

				return &lib.MidResponse{
					Data: result,
				}, nil
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
	routes := lib.NewRoutes().Get(lib.Route{
		Handler: lib.
			NewMiddleware().
			Build(func(r *http.Request) (*lib.MidResponse, error) {
				result := map[string]string{
					"message": "Hello, World!",
				}

				return &lib.MidResponse{
					Status: http.StatusOK,
					Data:   result,
				}, nil
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

func TestHttpRouterHandlerDuplicateSlashes(t *testing.T) {
	routes := lib.NewRoutes().Get(lib.Route{
		Handler: lib.NewMiddleware().
			Build(func(r *http.Request) (*lib.MidResponse, error) {
				result := map[string]string{"message": "Success"}

				return &lib.MidResponse{
					Status: http.StatusOK,
					Data:   result,
				}, nil
			}),
		Path: "/some/path",
	})

	// Petición con slashes duplicados
	req := httptest.NewRequest("GET", "http://example.com//some//path", nil)
	rec := httptest.NewRecorder()
	lib.HttpRouterHandler(routes.Routes)(rec, req)

	assert.EqualValues(t, http.StatusOK, rec.Code, "Debería normalizar los slashes y devolver 200 OK")

	var response map[string]string
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.EqualValues(t, "Success", response["message"])
}
