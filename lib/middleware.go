package lib

import (
	"context"
	"log/slog"
	"net/http"
)

type MiddlewareFunc func(w http.ResponseWriter, r *http.Request) (any, error)

type middleware struct {
	Ctx              context.Context
	MiddleFuncs      []MiddlewareFunc
	MiddleAfterFuncs []MiddlewareFunc
}

type Middleware interface {
	Use(MiddlewareFunc) Middleware
	After(MiddlewareFunc) Middleware
	Handler(fn MiddlewareFunc) MiddlewareFunc
}

func NewMiddleware(ctx context.Context) Middleware {
	middleFuncs := []MiddlewareFunc{}
	return &middleware{
		Ctx:              ctx,
		MiddleFuncs:      middleFuncs,
		MiddleAfterFuncs: []MiddlewareFunc{},
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

func (m *middleware) Handler(fn MiddlewareFunc) MiddlewareFunc {
	return func(w http.ResponseWriter, r *http.Request) (any, error) {
		// Every middleware before is executed
		for _, middleFunc := range m.MiddleFuncs {
			result, err := middleFunc(w, r)
			if err != nil {
				slog.Error("Error in middleware", slog.String("error", err.Error()))
				return nil, err
			}
			if result != nil {
				slog.Info("Middleware returned a result", slog.Any("result", result))
			}
		}

		result, err := fn(w, r)
		if err != nil {
			slog.Error("Error in handler", slog.String("error", err.Error()))
			return nil, err
		}
		if result != nil {
			slog.Info("Handler returned a result", slog.Any("result", result))
		}

		// Every middleware after is executed
		for _, middleAfterFunc := range m.MiddleAfterFuncs {
			result, err := middleAfterFunc(w, r)
			if err != nil {
				slog.Error("Error in middleware", slog.String("error", err.Error()))
				return nil, err
			}
			if result != nil {
				slog.Info("Middleware returned a result", slog.Any("result", result))
			}
		}
		return result, err
	}
}
