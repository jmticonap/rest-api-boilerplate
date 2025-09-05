package lib

import (
	"context"
	"log/slog"
	"net/http"
)

type MiddlewareFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) (any, error)

type middleware struct {
	Ctx              context.Context
	Response         http.ResponseWriter
	Request          *http.Request
	MiddleFuncs      []MiddlewareFunc
	MiddleAfterFuncs []MiddlewareFunc
}

type Middleware interface {
	Use(MiddlewareFunc) Middleware
	After(MiddlewareFunc) Middleware
	Handler(MiddlewareFunc) Middleware
}

func NewMiddleware(w http.ResponseWriter, r *http.Request) Middleware {
	middleFuncs := []MiddlewareFunc{}
	return &middleware{
		Ctx:         context.Background(),
		Response:    w,
		Request:     r,
		MiddleFuncs: middleFuncs,
	}
}

func (m *middleware) Use(fn MiddlewareFunc) Middleware {
	m.MiddleFuncs = append(m.MiddleFuncs, fn)

	return m
}

func (m *middleware) After(fn MiddlewareFunc) Middleware {
	m.MiddleAfterFuncs = append(m.MiddleAfterFuncs, fn)

	return m
}

func (m *middleware) Handler(fn MiddlewareFunc) Middleware {
	// Every middleware before is executed
	for _, middleFunc := range m.MiddleFuncs {
		result, err := middleFunc(m.Ctx, m.Response, m.Request)
		if err != nil {
			slog.Error("Error in middleware", slog.String("error", err.Error()))
			return m
		}
		if result != nil {
			slog.Info("Middleware returned a result", slog.Any("result", result))
		}
	}

	result, err := fn(m.Ctx, m.Response, m.Request)
	if err != nil {
		slog.Error("Error in handler", slog.String("error", err.Error()))
		return m
	}
	if result != nil {
		slog.Info("Handler returned a result", slog.Any("result", result))
	}

	// Every middleware after is executed
	for _, middleAfterFunc := range m.MiddleAfterFuncs {
		result, err := middleAfterFunc(m.Ctx, m.Response, m.Request)
		if err != nil {
			slog.Error("Error in middleware", slog.String("error", err.Error()))
			return m
		}
		if result != nil {
			slog.Info("Middleware returned a result", slog.Any("result", result))
		}
	}

	return m
}
