package handler

import (
	"errors"
	"net/http"
)

type Middleware interface {
	CallMiddleware(handler http.HandlerFunc) http.HandlerFunc
}

func NewMiddleware() Middleware {
	return &middleware{}
}

type middleware struct{}

func (m *middleware) CallMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		m.enablingCORS(handler)

		if req.Method != http.MethodPost {
			DefaultErrorResponse(rw, errors.New("method not allowed"), http.StatusMethodNotAllowed)
			return
		}

		handler.ServeHTTP(rw, req)
	}
}

func (m *middleware) enablingCORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		rw.Header().Set("Content-Type", "application-json")

		handler.ServeHTTP(rw, req)
	}
}
