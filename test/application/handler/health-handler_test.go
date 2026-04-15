package handler_test

import (
	"net/http"
	"testing"

	"rest-api/app/application/handler"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	res, err := handler.HealthCheck(req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusOK, res.Status)
	
	data, ok := res.Data.(map[string]string)
	assert.True(t, ok)
	assert.Equal(t, "UP", data["status"])
}
