package lib

import (
	"fmt"
	"strings"
)

func GetHandler(route RouteSchema, path string) (GenericHandler, error) {
	cleanPath, _ := strings.CutPrefix(path, "/")
	pathString := strings.Split(cleanPath, "/")
	pathLen := len(pathString)

	var handler GenericHandler

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

func HttpRouter(event RequestEvent, routes map[string]RouteSchema) (any, error) {
	method := event.Method
	path := event.Path

	var handler func() (any, error)
	var err error

	handler, err = GetHandler(routes[rootNodeKey].Children[method], path)

	if err == nil {
		return handler()
	} else {
		return nil, err
	}
}
