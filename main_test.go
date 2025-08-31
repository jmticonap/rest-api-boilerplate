package main

import (
	"fmt"
	"router-schema/lib"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	t.Run("should have two children in 'alumno'", func(t *testing.T) {
		router := lib.NewRouter()
		router.Get(lib.Route{
			Handler: func() {
				fmt.Println("First handler")
			},
			Path: "/alumno/notas/codigo",
		}).Get(lib.Route{
			Handler: func() {
				fmt.Println("class room")
			},
			Path: "/alumno/class-room",
		})
		expected := 2
		actual := len(router.Routes["/"].Children["alumno"].Children)

		assert.Equal(t, expected, actual, "The length in 'alumno' route isn't 2 children")
	})
	router := lib.NewRouter()
	router.Get(lib.Route{
		Handler: func() {},
		Path:    "/test",
	})

	handler, err := lib.GetHandler(router.Routes["/"], "/test")

	if err != nil {
		t.Errorf("Expected nil error, got %s", err)
	}

	if handler == nil {
		t.Error("Expected handler, got nil")
	}
}
