package mockit

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
)

// Response represents a mock response.
type Response struct {
	Code    int
	Headers map[string]string
	Body    string
}

// Endpoint represents a mock endpoint.
type Endpoint struct {
	URL      string
	Method   string
	Response Response
}

// Config represents a mock configuration.
type Config struct {
	Endpoints []Endpoint
}

// NewRouter returns a new router based on the given configuration.
func NewRouter(config Config) *httptreemux.TreeMux {
	router := httptreemux.New()

	for _, endpoint := range config.Endpoints {
		setupMock(router, endpoint)
	}

	return router
}

func setupMock(router *httptreemux.TreeMux, endpoint Endpoint) {
	router.Handle(endpoint.Method, endpoint.URL, func(w http.ResponseWriter, r *http.Request, m map[string]string) {
		w.WriteHeader(endpoint.Response.Code)
		for k, v := range endpoint.Response.Headers {
			w.Header().Set(k, v)
		}

		w.Write([]byte(endpoint.Response.Body))
	})
}
