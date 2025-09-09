package lib_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"rest-api/app/lib"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddlewareBefore(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/wrong/path", nil)
	rec := httptest.NewRecorder()

	result, _ := lib.NewMiddleware(t.Context()).
		Use(func(w http.ResponseWriter, r *http.Request) (any, error) {
			result := map[string]string{
				"message": "Hello, World!",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(result); err != nil {
				return nil, err
			}

			return result, nil
		}).Handler(func(w http.ResponseWriter, r *http.Request) (any, error) {
		return nil, nil
	})(rec, req)

	var response map[string]string
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.EqualValues(t, "Hello, World!", response["message"])

	mapResponse, _ := result.(map[string]string)
	assert.EqualValues(t, "Hello, World!", mapResponse["message"])
}

func TestMiddlewareAfter(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/wrong/path", nil)
	rec := httptest.NewRecorder()

	lib.NewMiddleware(t.Context()).
		After(func(w http.ResponseWriter, r *http.Request) (any, error) {
			result := map[string]string{
				"message": "Hello, World!",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(result); err != nil {
				return nil, err
			}

			return result, nil
		}).Handler(func(w http.ResponseWriter, r *http.Request) (any, error) {
		return nil, nil
	})(rec, req)

	var response map[string]string
	json.Unmarshal(rec.Body.Bytes(), &response)
	assert.EqualValues(t, "Hello, World!", response["message"])
}

func TestMiddlewareError(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/wrong/path", nil)
	rec := httptest.NewRecorder()

	_, resError := lib.NewMiddleware(t.Context()).
		Handler(func(w http.ResponseWriter, r *http.Request) (any, error) {
			return nil, fmt.Errorf("Some error")
		})(rec, req)

	assert.EqualError(t, resError, "Some error")
}

func TestMiddlewareBeforeError(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/wrong/path", nil)
	rec := httptest.NewRecorder()

	_, resError := lib.NewMiddleware(t.Context()).
		Use(func(w http.ResponseWriter, r *http.Request) (any, error) {
			return nil, fmt.Errorf("Some error")
		}).Handler(func(w http.ResponseWriter, r *http.Request) (any, error) {
		return nil, nil
	})(rec, req)

	assert.EqualError(t, resError, "Some error")
}

func TestMiddlewareAfterError(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/wrong/path", nil)
	rec := httptest.NewRecorder()

	_, resError := lib.NewMiddleware(t.Context()).
		After(func(w http.ResponseWriter, r *http.Request) (any, error) {
			return nil, fmt.Errorf("Some error")
		}).Handler(func(w http.ResponseWriter, r *http.Request) (any, error) {
		return nil, nil
	})(rec, req)

	assert.EqualError(t, resError, "Some error")
}
