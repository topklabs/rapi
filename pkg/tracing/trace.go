package tracing

import "net/http"

// Tracer is a tracing enabled middleware.
type Tracer struct {
	Middleware func(handler http.Handler) http.Handler
}
