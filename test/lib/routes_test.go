package lib_test

import (
	"context"
	"fmt"
	"net/http"
	"router-schema/app/lib"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutesGet(t *testing.T) {
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
		cNotas, okNotas := cAlumno.Children["notas"]
		_, okCodigo := cNotas.Children["codigo"]
		_, okClassRoom := cAlumno.Children["class-room"]

		assert.True(t, okGet)
		assert.True(t, okAlumno)
		assert.True(t, okNotas)
		assert.True(t, okCodigo)
		assert.True(t, okClassRoom)
		assert.True(t, okRoot)
	})
}

func TestRoutesPost(t *testing.T) {
	ctx := context.Background()
	routes := lib.NewRoutes(nil).Post(lib.Route{
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
	}).Post(lib.Route{
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
		cGet, okGet := cRoot.Children[http.MethodPost]
		cAlumno, okAlumno := cGet.Children["alumno"]
		cNotas, okNotas := cAlumno.Children["notas"]
		_, okCodigo := cNotas.Children["codigo"]
		_, okClassRoom := cAlumno.Children["class-room"]

		assert.True(t, okGet)
		assert.True(t, okAlumno)
		assert.True(t, okNotas)
		assert.True(t, okCodigo)
		assert.True(t, okClassRoom)
		assert.True(t, okRoot)
	})
}

func TestRoutesPatch(t *testing.T) {
	ctx := context.Background()
	routes := lib.NewRoutes(nil).Patch(lib.Route{
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
	}).Patch(lib.Route{
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
		cGet, okGet := cRoot.Children[http.MethodPatch]
		cAlumno, okAlumno := cGet.Children["alumno"]
		cNotas, okNotas := cAlumno.Children["notas"]
		_, okCodigo := cNotas.Children["codigo"]
		_, okClassRoom := cAlumno.Children["class-room"]

		assert.True(t, okGet)
		assert.True(t, okAlumno)
		assert.True(t, okNotas)
		assert.True(t, okCodigo)
		assert.True(t, okClassRoom)
		assert.True(t, okRoot)
	})
}

func TestRoutesPut(t *testing.T) {
	ctx := context.Background()
	routes := lib.NewRoutes(nil).Put(lib.Route{
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
	}).Put(lib.Route{
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
		cGet, okGet := cRoot.Children[http.MethodPut]
		cAlumno, okAlumno := cGet.Children["alumno"]
		cNotas, okNotas := cAlumno.Children["notas"]
		_, okCodigo := cNotas.Children["codigo"]
		_, okClassRoom := cAlumno.Children["class-room"]

		assert.True(t, okGet)
		assert.True(t, okAlumno)
		assert.True(t, okNotas)
		assert.True(t, okCodigo)
		assert.True(t, okClassRoom)
		assert.True(t, okRoot)
	})
}

func TestRoutesDelete(t *testing.T) {
	ctx := context.Background()
	routes := lib.NewRoutes(nil).Delete(lib.Route{
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
	}).Delete(lib.Route{
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
		cGet, okGet := cRoot.Children[http.MethodDelete]
		cAlumno, okAlumno := cGet.Children["alumno"]
		cNotas, okNotas := cAlumno.Children["notas"]
		_, okCodigo := cNotas.Children["codigo"]
		_, okClassRoom := cAlumno.Children["class-room"]

		assert.True(t, okGet)
		assert.True(t, okAlumno)
		assert.True(t, okNotas)
		assert.True(t, okCodigo)
		assert.True(t, okClassRoom)
		assert.True(t, okRoot)
	})
}

func TestRoutesConnect(t *testing.T) {
	ctx := context.Background()
	routes := lib.NewRoutes(nil).Connect(lib.Route{
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
	}).Connect(lib.Route{
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
		cGet, okGet := cRoot.Children[http.MethodConnect]
		cAlumno, okAlumno := cGet.Children["alumno"]
		cNotas, okNotas := cAlumno.Children["notas"]
		_, okCodigo := cNotas.Children["codigo"]
		_, okClassRoom := cAlumno.Children["class-room"]

		assert.True(t, okGet)
		assert.True(t, okAlumno)
		assert.True(t, okNotas)
		assert.True(t, okCodigo)
		assert.True(t, okClassRoom)
		assert.True(t, okRoot)
	})
}

func TestRoutesHead(t *testing.T) {
	ctx := context.Background()
	routes := lib.NewRoutes(nil).Head(lib.Route{
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
	}).Head(lib.Route{
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
		cGet, okGet := cRoot.Children[http.MethodHead]
		cAlumno, okAlumno := cGet.Children["alumno"]
		cNotas, okNotas := cAlumno.Children["notas"]
		_, okCodigo := cNotas.Children["codigo"]
		_, okClassRoom := cAlumno.Children["class-room"]

		assert.True(t, okGet)
		assert.True(t, okAlumno)
		assert.True(t, okNotas)
		assert.True(t, okCodigo)
		assert.True(t, okClassRoom)
		assert.True(t, okRoot)
	})
}

func TestRoutesOptions(t *testing.T) {
	ctx := context.Background()
	routes := lib.NewRoutes(nil).Options(lib.Route{
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
	}).Options(lib.Route{
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
		cGet, okGet := cRoot.Children[http.MethodOptions]
		cAlumno, okAlumno := cGet.Children["alumno"]
		cNotas, okNotas := cAlumno.Children["notas"]
		_, okCodigo := cNotas.Children["codigo"]
		_, okClassRoom := cAlumno.Children["class-room"]

		assert.True(t, okGet)
		assert.True(t, okAlumno)
		assert.True(t, okNotas)
		assert.True(t, okCodigo)
		assert.True(t, okClassRoom)
		assert.True(t, okRoot)
	})
}

func TestRoutesTrace(t *testing.T) {
	ctx := context.Background()
	routes := lib.NewRoutes(nil).Trace(lib.Route{
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
	}).Trace(lib.Route{
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
		cGet, okGet := cRoot.Children[http.MethodTrace]
		cAlumno, okAlumno := cGet.Children["alumno"]
		cNotas, okNotas := cAlumno.Children["notas"]
		_, okCodigo := cNotas.Children["codigo"]
		_, okClassRoom := cAlumno.Children["class-room"]

		assert.True(t, okGet)
		assert.True(t, okAlumno)
		assert.True(t, okNotas)
		assert.True(t, okCodigo)
		assert.True(t, okClassRoom)
		assert.True(t, okRoot)
	})
}

func TestRoutesMix(t *testing.T) {
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
	}).Post(lib.Route{
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
		cNotas, okNotas := cAlumno.Children["notas"]
		_, okCodigo := cNotas.Children["codigo"]
		cPost, okPost := cRoot.Children[http.MethodPost]
		cPostAlumno, okPostAlumno := cPost.Children["alumno"]
		_, okClassRoom := cPostAlumno.Children["class-room"]

		assert.True(t, okRoot)
		assert.True(t, okGet)
		assert.True(t, okAlumno)
		assert.True(t, okNotas)
		assert.True(t, okCodigo)
		assert.True(t, okPost)
		assert.True(t, okPostAlumno)
		assert.True(t, okClassRoom)
	})
}

func TestRoutesCleanRoutes(t *testing.T) {
	routes := lib.NewRoutes(nil)

	t.Run("Should haves all http method", func(t *testing.T) {
		methodsList := []string{
			http.MethodConnect,
			http.MethodDelete,
			http.MethodGet,
			http.MethodHead,
			http.MethodOptions,
			http.MethodPatch,
			http.MethodPost,
			http.MethodPut,
			http.MethodTrace,
		}
		cRoot, okRoot := routes.Routes["root"]

		assert.True(t, okRoot)
		for _, method := range methodsList {
			_, ok := cRoot.Children[method]

			assert.True(t, ok, "Method %s not found", method)
		}
	})
}
