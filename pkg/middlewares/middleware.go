package middlewares

import (
	"net/http"

	"github.com/rs/zerolog"
)

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn(w, r)
}

type Middleware struct {
	logger *zerolog.Logger
}

func NewMiddleware(logger *zerolog.Logger) *Middleware {
	return &Middleware{logger: logger}
}

func Chain(funcHandler Handler, middleware ...*Middleware) Handler {
	for _, m := range middleware {
		funcHandler = m.Handle(funcHandler)
	}
	return funcHandler
}
