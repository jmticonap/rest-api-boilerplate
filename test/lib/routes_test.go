package lib_test

import (
	"fmt"
	"net/http"
	"rest-api/app/lib"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutesGet(t *testing.T) {
	routes := lib.NewRoutes().
		Get(lib.Route{
			Handler: lib.NewMiddleware().
				Use(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware before handler executed")
					return nil, nil
				}).
				After(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware after handler executed")
					return nil, nil
				}).
				Build(func(r *http.Request) (*lib.MidResponse, error) {
					return nil, nil
				}),
			Path: "/alumno/notas/codigo",
		}).
		Get(lib.Route{
			Handler: lib.NewMiddleware().
				Build(func(r *http.Request) (*lib.MidResponse, error) {
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
	routes := lib.NewRoutes().
		Post(lib.Route{
			Handler: lib.NewMiddleware().
				Use(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware before handler executed")
					return nil, nil
				}).
				After(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware after handler executed")
					return nil, nil
				}).
				Build(func(r *http.Request) (*lib.MidResponse, error) {
					return nil, nil
				}),
			Path: "/alumno/notas/codigo",
		}).
		Post(lib.Route{
			Handler: lib.NewMiddleware().
				Build(func(r *http.Request) (*lib.MidResponse, error) {
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
	routes := lib.NewRoutes().
		Patch(lib.Route{
			Handler: lib.
				NewMiddleware().
				Use(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware before handler executed")
					return nil, nil
				}).
				After(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware after handler executed")
					return nil, nil
				}).
				Build(func(r *http.Request) (*lib.MidResponse, error) {
					return nil, nil
				}),
			Path: "/alumno/notas/codigo",
		}).
		Patch(lib.Route{
			Handler: lib.NewMiddleware().
				Build(func(r *http.Request) (*lib.MidResponse, error) {
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
	routes := lib.NewRoutes().
		Put(lib.Route{
			Handler: lib.
				NewMiddleware().
				Use(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware before handler executed")
					return nil, nil
				}).
				After(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware after handler executed")
					return nil, nil
				}).
				Build(func(r *http.Request) (*lib.MidResponse, error) {
					return nil, nil
				}),
			Path: "/alumno/notas/codigo",
		}).
		Put(lib.Route{
			Handler: lib.NewMiddleware().
				Build(func(r *http.Request) (*lib.MidResponse, error) {
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
	routes := lib.NewRoutes().
		Delete(lib.Route{
			Handler: lib.NewMiddleware().
				Use(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware before handler executed")
					return nil, nil
				}).
				After(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware after handler executed")
					return nil, nil
				}).
				Build(func(r *http.Request) (*lib.MidResponse, error) {
					return nil, nil
				}),
			Path: "/alumno/notas/codigo",
		}).
		Delete(lib.Route{
			Handler: lib.NewMiddleware().
				Build(func(r *http.Request) (*lib.MidResponse, error) {
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
	routes := lib.NewRoutes().
		Connect(lib.Route{
			Handler: lib.NewMiddleware().
				Use(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware before handler executed")
					return nil, nil
				}).
				After(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware after handler executed")
					return nil, nil
				}).
				Build(func(r *http.Request) (*lib.MidResponse, error) {
					return nil, nil
				}),
			Path: "/alumno/notas/codigo",
		}).
		Connect(lib.Route{
			Handler: lib.NewMiddleware().
				Build(func(r *http.Request) (*lib.MidResponse, error) {
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
	routes := lib.NewRoutes().
		Head(lib.Route{
			Handler: lib.NewMiddleware().
				Use(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware before handler executed")
					return nil, nil
				}).
				After(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware after handler executed")
					return nil, nil
				}).
				Build(func(r *http.Request) (*lib.MidResponse, error) {
					return nil, nil
				}),
			Path: "/alumno/notas/codigo",
		}).
		Head(lib.Route{
			Handler: lib.NewMiddleware().
				Build(func(r *http.Request) (*lib.MidResponse, error) {
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
	routes := lib.NewRoutes().
		Options(lib.Route{
			Handler: lib.NewMiddleware().
				Use(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware before handler executed")
					return nil, nil
				}).
				After(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware after handler executed")
					return nil, nil
				}).
				Build(func(r *http.Request) (*lib.MidResponse, error) {
					return nil, nil
				}),
			Path: "/alumno/notas/codigo",
		}).
		Options(lib.Route{
			Handler: lib.NewMiddleware().
				Build(func(r *http.Request) (*lib.MidResponse, error) {
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
	routes := lib.NewRoutes().
		Trace(lib.Route{
			Handler: lib.NewMiddleware().
				Use(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware before handler executed")
					return nil, nil
				}).
				After(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware after handler executed")
					return nil, nil
				}).
				Build(func(r *http.Request) (*lib.MidResponse, error) {
					return nil, nil
				}),
			Path: "/alumno/notas/codigo",
		}).
		Trace(lib.Route{
			Handler: lib.NewMiddleware().
				Build(func(r *http.Request) (*lib.MidResponse, error) {
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
	routes := lib.NewRoutes().
		Get(lib.Route{
			Handler: lib.NewMiddleware().
				Use(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware before handler executed")
					return nil, nil
				}).
				After(func(r *http.Request) (*lib.MidResponse, error) {
					fmt.Println("Middleware after handler executed")
					return nil, nil
				}).
				Build(func(r *http.Request) (*lib.MidResponse, error) {
					return nil, nil
				}),
			Path: "/alumno/notas/codigo",
		}).
		Post(lib.Route{
			Handler: lib.NewMiddleware().
				Build(func(r *http.Request) (*lib.MidResponse, error) {
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
	routes := lib.NewRoutes()

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

func TestRouteIntegrityOverwriting(t *testing.T) {
	routes := lib.NewRoutes()

	// 1. Registramos primero la ruta larga
	routes.Get(lib.Route{
		Handler: lib.NewMiddleware().
			Build(func(r *http.Request) (*lib.MidResponse, error) {
				return nil, nil
			}),
		Path: "/alumno/notas/codigo",
	})

	// 2. Registramos después la ruta corta (el padre)
	routes.Get(lib.Route{
		Handler: lib.NewMiddleware().
			Build(func(r *http.Request) (*lib.MidResponse, error) {
				return nil, nil
			}),
		Path: "/alumno/notas",
	})

	t.Run("Should not delete child routes when a parent is registered", func(t *testing.T) {
		cRoot := routes.Routes["root"]
		cGet := cRoot.Children[http.MethodGet]
		cAlumno := cGet.Children["alumno"]
		cNotas := cAlumno.Children["notas"]

		// Verificamos si "codigo" sigue existiendo
		_, okCodigo := cNotas.Children["codigo"]

		assert.True(t, okCodigo, "CRÍTICO: La ruta /alumno/notas/codigo DESAPARECIÓ al registrar /alumno/notas")
	})
}
