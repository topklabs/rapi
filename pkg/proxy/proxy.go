package proxy

import (
	"context"
	"net/http"
)

// Proxy is a reverse proxy for passing requests through to an external API.
type Proxy struct {
	context     context.Context
	config      Config
	middlewares []Middleware
}

// NewProxy instantiates a new instance of the proxy server.
func NewProxy(ctx context.Context, config Config) Proxy {
	return Proxy{
		context:     ctx,
		config:      config,
		middlewares: []Middleware{},
	}
}

// ServeHTTP implements a Go HTTP Handler interface
func (p Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	apply(http.HandlerFunc(p.proxyRequest), p.middlewares).ServeHTTP(w, r.WithContext(p.context))
}

func (p Proxy) proxyRequest(w http.ResponseWriter, r *http.Request) { p.ServeHTTP(w, r) }

func apply(handler http.Handler, middlewares []Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
