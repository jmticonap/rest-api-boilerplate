package lib

import (
	"fmt"
	"net/http"
	"strings"
)

type RequestEvent struct {
	Method  string
	Path    string
	Headers map[string]string
	Body    map[string]any
}

type Route struct {
	Method  string
	Path    string
	Handler func()
}

type RouteSchema struct {
	Handler  func()
	Children map[string]RouteSchema
}

type router struct {
	Routes map[string]RouteSchema
}

type Router interface {
	Connect(route Route) Router
	Delete(route Route) Router
	Get(route Route) Router
	Head(route Route) Router
	Options(route Route) Router
	Patch(route Route) Router
	Post(route Route) Router
	Put(route Route) Router
	Trace(route Route) Router
}

var rootNodeKey string = "root"

func NewRouter() *router {
	routes := make(map[string]RouteSchema)
	routes[rootNodeKey] = RouteSchema{
		Handler:  nil,
		Children: make(map[string]RouteSchema),
	}
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
	for _, method := range methodsList {
		rSch := RouteSchema{
			Handler:  nil,
			Children: make(map[string]RouteSchema),
		}
		routes[rootNodeKey].Children[method] = rSch
	}

	return &router{
		Routes: routes,
	}
}

func (r *router) Connect(route Route) Router {
	AddRoute(r.Routes[rootNodeKey].Children[http.MethodConnect], route)
	return r
}
func (r *router) Delete(route Route) Router {
	AddRoute(r.Routes[rootNodeKey].Children[http.MethodDelete], route)
	return r
}
func (r *router) Get(route Route) Router {
	AddRoute(r.Routes[rootNodeKey].Children[http.MethodGet], route)
	return r
}
func (r *router) Head(route Route) Router {
	AddRoute(r.Routes[rootNodeKey].Children[http.MethodHead], route)
	return r
}
func (r *router) Options(route Route) Router {
	AddRoute(r.Routes[rootNodeKey].Children[http.MethodOptions], route)
	return r
}
func (r *router) Patch(route Route) Router {
	AddRoute(r.Routes[rootNodeKey].Children[http.MethodPatch], route)
	return r
}
func (r *router) Post(route Route) Router {
	AddRoute(r.Routes[rootNodeKey].Children[http.MethodPost], route)
	return r
}
func (r *router) Put(route Route) Router {
	AddRoute(r.Routes[rootNodeKey].Children[http.MethodPut], route)
	return r
}
func (r *router) Trace(route Route) Router {
	AddRoute(r.Routes[rootNodeKey].Children[http.MethodTrace], route)
	return r
}

func AddRoute(rootRoute RouteSchema, route Route) {
	cleanPath, _ := strings.CutPrefix(route.Path, "/")
	pathString := strings.Split(cleanPath, "/")
	pathLen := len(pathString)

	if pathLen == 1 {
		rootRoute.Children[pathString[0]] = RouteSchema{
			Handler:  route.Handler,
			Children: make(map[string]RouteSchema),
		}
	} else {
		if _, ok := rootRoute.Children[pathString[0]]; !ok {
			rootRoute.Children[pathString[0]] = RouteSchema{
				Handler:  nil,
				Children: make(map[string]RouteSchema),
			}
		}

		AddRoute(rootRoute.Children[pathString[0]], Route{
			Handler: route.Handler,
			Path:    strings.Join(pathString[1:], "/"),
		})
	}
}

func GetHandler(route RouteSchema, path string) (func(), error) {
	cleanPath, _ := strings.CutPrefix(path, "/")
	pathString := strings.Split(cleanPath, "/")
	pathLen := len(pathString)

	var handler func()

	if rt, ok := route.Children[pathString[0]]; ok {
		if pathLen == 1 {
			handler = route.Children[pathString[0]].Handler
		} else {
			return GetHandler(
				rt, strings.Join(pathString[1:], "/"),
			)
		}
	} else {
		return nil, fmt.Errorf("Path not found\n")
	}

	return handler, nil
}

func ExeEvent(event RequestEvent, routes map[string]RouteSchema) {
	method := event.Method
	path := event.Path

	var handler func()
	var err error

	handler, err = GetHandler(routes[rootNodeKey].Children[method], path)

	if err == nil {
		handler()
	} else {
		fmt.Printf("%s", err)
	}
}
