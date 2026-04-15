package lib_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"rest-api/app/lib"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddlewareBeforeResponse(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/wrong/path", nil)
	rec := httptest.NewRecorder()
	isHandlerExecute := false

	res, err := lib.NewMiddleware().
		Use(func(r *http.Request) (*lib.MidResponse, error) {
			result := map[string]string{
				"message": "Hello, World!",
			}

			return &lib.MidResponse{
				Status: http.StatusOK,
				Data:   result,
			}, nil
		}).
		Build(func(r *http.Request) (*lib.MidResponse, error) {
			isHandlerExecute = true
			return nil, nil
		})(&lib.ResponseWrapper{ResponseWriter: rec}, req)

	result, ok := res.Data.(map[string]string)

	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, res.Data, result)
	assert.False(t, isHandlerExecute)
}

func TestMiddlewareAfter(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/wrong/path", nil)
	rec := httptest.NewRecorder()
	isAfterExecute := false

	res, err := lib.NewMiddleware().
		Use(func(r *http.Request) (*lib.MidResponse, error) {
			result := map[string]string{
				"message": "Hello, World!",
			}

			return &lib.MidResponse{
				Status: http.StatusOK,
				Data:   result,
			}, nil
		}).
		After(func(r *http.Request) (*lib.MidResponse, error) {
			isAfterExecute = true
			return nil, nil
		}).
		Build(func(r *http.Request) (*lib.MidResponse, error) {
			return nil, nil
		})(&lib.ResponseWrapper{ResponseWriter: rec}, req)

	result, ok := res.Data.(map[string]string)

	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, res.Data, result)
	assert.True(t, isAfterExecute)
}

func TestMiddlewareError(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/wrong/path", nil)
	rec := httptest.NewRecorder()
	isHandlerExecute := false

	res, err := lib.NewMiddleware().
		Use(func(r *http.Request) (*lib.MidResponse, error) {
			return nil, fmt.Errorf("test error")
		}).
		Error(func(r *http.Request) (*lib.MidResponse, error) {
			return &lib.MidResponse{
				Status: http.StatusBadRequest,
				Data:   nil,
			}, nil
		}).
		Build(func(r *http.Request) (*lib.MidResponse, error) {
			isHandlerExecute = true
			return nil, nil
		})(&lib.ResponseWrapper{ResponseWriter: rec}, req)

	assert.NoError(t, err)
	assert.Equal(t, res.Status, http.StatusBadRequest)
	assert.False(t, isHandlerExecute)
}

func TestMiddlewareErrorFail(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/wrong/path", nil)
	rec := httptest.NewRecorder()
	isHandlerExecute := false

	res, err := lib.NewMiddleware().
		Use(func(r *http.Request) (*lib.MidResponse, error) {
			return nil, fmt.Errorf("test error")
		}).
		Error(func(r *http.Request) (*lib.MidResponse, error) {
			return &lib.MidResponse{
				Status: http.StatusBadRequest,
				Data:   nil,
			}, nil
		}).
		Error(func(r *http.Request) (*lib.MidResponse, error) {
			return nil, fmt.Errorf("test error fail")
		}).
		Build(func(r *http.Request) (*lib.MidResponse, error) {
			isHandlerExecute = true
			return nil, nil
		})(&lib.ResponseWrapper{ResponseWriter: rec}, req)

	assert.Error(t, err)
	assert.Nil(t, res)
	assert.False(t, isHandlerExecute)
}

func TestMiddlewareBeforeError(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/wrong/path", nil)
	rec := httptest.NewRecorder()

	_, err := lib.NewMiddleware().
		Use(func(r *http.Request) (*lib.MidResponse, error) {
			return nil, fmt.Errorf("Some error")
		}).
		Build(func(r *http.Request) (*lib.MidResponse, error) {
			return nil, nil
		})(&lib.ResponseWrapper{ResponseWriter: rec}, req)

	assert.EqualError(t, err, "Some error")
}

func TestMiddlewareAfterError(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/wrong/path", nil)
	rec := httptest.NewRecorder()

	afterExecuted := false

	_, err := lib.NewMiddleware().
		After(func(r *http.Request) (*lib.MidResponse, error) {
			afterExecuted = true
			return nil, fmt.Errorf("After error")
		}).
		Build(func(r *http.Request) (*lib.MidResponse, error) {
			return nil, fmt.Errorf("Main error")
		})(&lib.ResponseWrapper{ResponseWriter: rec}, req)

	assert.EqualError(t, err, "Main error")
	assert.True(t, afterExecuted)
}

func TestMiddlewareFlowBypassVulnerability(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/sensitive-data", nil)
	rec := httptest.NewRecorder()

	handlerExecuted := false

	_, err := lib.NewMiddleware().
		Use(func(r *http.Request) (*lib.MidResponse, error) {
			return &lib.MidResponse{
				Data: "Unauthorized",
			}, fmt.Errorf("Unauthorized")
		}).
		Use(func(r *http.Request) (*lib.MidResponse, error) {
			return nil, nil
		}).
		Build(func(r *http.Request) (*lib.MidResponse, error) {
			handlerExecuted = true
			return &lib.MidResponse{Data: "Sensitive Data"}, nil
		})(&lib.ResponseWrapper{ResponseWriter: rec}, req)

	assert.Error(t, err)
	assert.False(t, handlerExecuted, "VULNERABILIDAD: El handler se ejecutó a pesar de que el primer middleware falló")
}

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
