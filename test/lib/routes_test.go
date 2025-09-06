package lib_test

import (
	"context"
	"fmt"
	"net/http"
	"router-schema/app/lib"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	ctx := context.Background()
	routes := lib.NewRoutes(nil).Get(lib.Route{
		Handler: lib.
			NewMiddleware(ctx).
			Use(func(w http.ResponseWriter, r *http.Request) (any, error) {
				fmt.Println("Middleware before handler executed")
				return nil, nil
			}).
			After(func(w http.ResponseWriter, r *http.Request) (any, error) {
				fmt.Println("Middleware after handler executed")
				return nil, nil
			}).
			Handler(func(w http.ResponseWriter, r *http.Request) (any, error) {
				return nil, nil
			}),
		Path: "/alumno/notas/codigo",
	}).Get(lib.Route{
		Handler: lib.NewMiddleware(ctx).
			Handler(func(w http.ResponseWriter, r *http.Request) (any, error) {
				fmt.Println("Hello from middleware")
				fmt.Println("class room")
				return nil, nil
			}),
		Path: "/alumno/class-room",
	})

	t.Run("Should haves the expected keys", func(t *testing.T) {
		cRoot, okRoot := routes.Routes["root"]
		cGet, okGet := cRoot.Children[http.MethodGet]
		cAlumno, okAlumno := cGet.Children["alumno"]
		_, okNotas := cAlumno.Children["notas"]
		_, okCodigo := cAlumno.Children["notas"].Children["codigo"]
		_, okClassRoom := cAlumno.Children["class-room"]

		assert.True(t, okGet)
		assert.True(t, okAlumno)
		assert.True(t, okNotas)
		assert.True(t, okCodigo)
		assert.True(t, okClassRoom)
		assert.True(t, okRoot)
	})
}
