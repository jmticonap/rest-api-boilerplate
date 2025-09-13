package lib

import (
	"context"
	"log/slog"
	"net/http"
)

type MiddlewareFunc func(w http.ResponseWriter, r *http.Request) (any, error)

type Middleware struct {
	Ctx              context.Context
	MiddleFuncs      []MiddlewareFunc
	MiddleAfterFuncs []MiddlewareFunc
}

type IMiddleware interface {
	Use(MiddlewareFunc) Middleware
	After(MiddlewareFunc) Middleware
	Handler(fn MiddlewareFunc) MiddlewareFunc
}

func NewMiddleware(ctx context.Context) *Middleware {
	return &Middleware{
		Ctx:              ctx,
		MiddleFuncs:      []MiddlewareFunc{},
		MiddleAfterFuncs: []MiddlewareFunc{},
	}
}

func (m *Middleware) Use(fn MiddlewareFunc) *Middleware {
	m.MiddleFuncs = append(m.MiddleFuncs, fn)
	return m
}

func (m *Middleware) After(fn MiddlewareFunc) *Middleware {
	m.MiddleAfterFuncs = append(m.MiddleAfterFuncs, fn)
	return m
}

func (m *Middleware) Handler(fn MiddlewareFunc) MiddlewareFunc {
	return func(w http.ResponseWriter, r *http.Request) (any, error) {
		var resultBf any
		// Every middleware before is executed
		for _, middleFunc := range m.MiddleFuncs {
			var err error
			resultBf, err = middleFunc(w, r)
			if err != nil {
				slog.Error("Error in middleware", slog.String("error", err.Error()))
				return nil, err
			}
			if resultBf != nil {
				slog.Info("Middleware returned a result", slog.Any("result", resultBf))
				break
			}
		}
		if resultBf != nil {
			return resultBf, nil
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
