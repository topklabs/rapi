package proxy

import "net/http"

// Middleware describes the signature of a middleware function that can be applied to the proxy handler.
type Middleware func(http.Handler) http.Handler
