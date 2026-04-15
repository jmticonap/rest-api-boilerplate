package handler_test

import (
	"net/http/httptest"
	"rest-api/app/application/handler"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreetingIndex(t *testing.T) {
	t.Run("Exec GreetingIndex", func(t *testing.T) {
		// rec := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "http://example.com/some/path", nil)
		response, err := handler.GreetingIndex(req)

		assert.Nil(t, response)
		assert.Nil(t, err)
	})
}
