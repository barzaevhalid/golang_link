package middleware

import "net/http"

type MiddleWare func(http.Handler) http.Handler

func Chain(middlewares ...MiddleWare) MiddleWare {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}
