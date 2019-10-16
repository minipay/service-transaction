package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"service-transaction/utils/errorcollector"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct{}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	next(w, r)
}

// PanicCatcher is use for collecting panic that happened in endpoint
func (m *GoMiddleware) PanicCatcher(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		rec := recover()
		if rec != nil {
			errPanic := fmt.Sprintf("Endpoint: %s - panic: %v", r.URL.RequestURI(), rec)
			errorcollector.WritePanic(errPanic, debug.Stack())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}()
	next(w, r)
}

// InitMiddleware intialize the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
