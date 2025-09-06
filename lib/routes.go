package lib

import (
	"net/http"
	"strings"
)

type GenericHandler func() (any, error)

type RequestEvent struct {
	Method  string
	Path    string
	Headers map[string]string
	Body    map[string]any
}

type Route struct {
	Method  string
	Path    string
	Handler MiddlewareFunc
}

type RouteSchema struct {
	Handler  MiddlewareFunc
	Children map[string]RouteSchema
}

type Routes struct {
	Routes map[string]RouteSchema
}

type IRoutes interface {
	Connect(route Route) IRoutes
	Delete(route Route) IRoutes
	Get(route Route) IRoutes
	Head(route Route) IRoutes
	Options(route Route) IRoutes
	Patch(route Route) IRoutes
	Post(route Route) IRoutes
	Put(route Route) IRoutes
	Trace(route Route) IRoutes
}

var rootNodeKey string = "root"

func GetRouteList() map[string]RouteSchema {
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

	return routes
}

func NewRoutes(routeList map[string]RouteSchema) *Routes {
	if routeList == nil {
		routeList = GetRouteList()
	}

	return &Routes{
		Routes: routeList,
	}
}

func (r *Routes) Connect(route Route) IRoutes {
	AddRoute(r.Routes[rootNodeKey].Children[http.MethodConnect], route)
	return r
}
func (r *Routes) Delete(route Route) IRoutes {
	AddRoute(r.Routes[rootNodeKey].Children[http.MethodDelete], route)
	return r
}
func (r *Routes) Get(route Route) IRoutes {
	AddRoute(r.Routes[rootNodeKey].Children[http.MethodGet], route)
	return r
}
func (r *Routes) Head(route Route) IRoutes {
	AddRoute(r.Routes[rootNodeKey].Children[http.MethodHead], route)
	return r
}
func (r *Routes) Options(route Route) IRoutes {
	AddRoute(r.Routes[rootNodeKey].Children[http.MethodOptions], route)
	return r
}
func (r *Routes) Patch(route Route) IRoutes {
	AddRoute(r.Routes[rootNodeKey].Children[http.MethodPatch], route)
	return r
}
func (r *Routes) Post(route Route) IRoutes {
	AddRoute(r.Routes[rootNodeKey].Children[http.MethodPost], route)
	return r
}
func (r *Routes) Put(route Route) IRoutes {
	AddRoute(r.Routes[rootNodeKey].Children[http.MethodPut], route)
	return r
}
func (r *Routes) Trace(route Route) IRoutes {
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
