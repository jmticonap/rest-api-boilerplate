package lib_test

import (
	"net/http"
	"net/http/httptest"
	"rest-api/app/lib"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleware_ShouldStopIfBeforeReturnsResponse(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	handlerExecuted := false
	afterExecuted := false

	m := lib.NewMiddleware().
		Use(func(r *http.Request) (*lib.MidResponse, error) {
			// Returns a response with Status and Data
			return &lib.MidResponse{
				Status: http.StatusForbidden,
				Data:   map[string]string{"error": "stopped"},
			}, nil
		}).
		After(func(r *http.Request) (*lib.MidResponse, error) {
			afterExecuted = true
			return nil, nil
		})
		// func(ctx context.Context, r *http.Request) (*MidResponse, error)
	handler := m.Build(func(r *http.Request) (*lib.MidResponse, error) {
		handlerExecuted = true
		return nil, nil
	})

	_, err := handler(&lib.ResponseWrapper{ResponseWriter: rec}, req)

	assert.NoError(t, err)
	assert.False(t, handlerExecuted, "Handler should NOT have been executed")
	assert.True(t, afterExecuted, "After middleware SHOULD have been executed")
	assert.Equal(t, http.StatusForbidden, rec.Code)
	assert.Contains(t, rec.Body.String(), "stopped")
}

func TestMiddleware_ShouldAlwaysRunAfter(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	afterExecuted := false

	m := lib.NewMiddleware().
		Use(func(r *http.Request) (*lib.MidResponse, error) {
			return nil, nil
		}).
		After(func(r *http.Request) (*lib.MidResponse, error) {
			afterExecuted = true
			return nil, nil
		})

	handler := m.Build(func(r *http.Request) (*lib.MidResponse, error) {
		return nil, nil
	})

	handler(&lib.ResponseWrapper{ResponseWriter: rec}, req)

	assert.True(t, afterExecuted, "After middleware SHOULD have been executed")
}
