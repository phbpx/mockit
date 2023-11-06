package mockit

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dimfeld/httptreemux"
)

// regex to match {{ param }}
var paramRegex = regexp.MustCompile(`{{.+}}`)

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
func NewRouter(config Config) http.Handler {
	router := httptreemux.New()

	for _, e := range config.Endpoints {
		endpoint := e

		router.Handle(endpoint.Method, endpoint.URL, func(w http.ResponseWriter, r *http.Request, params map[string]string) {
			body, err := compile(endpoint.Response.Body, params)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			for k, v := range endpoint.Response.Headers {
				w.Header().Set(k, v)
			}

			w.WriteHeader(endpoint.Response.Code)
			if _, err := w.Write(body); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
	}

	return router
}

func compile(text string, params map[string]string) ([]byte, error) {
	if !paramRegex.MatchString(text) {
		return []byte(text), nil
	}

	funcMap := template.FuncMap{
		"urlParam":   paramsValue(params),
		"uuid":       gofakeit.UUID,
		"now":        time.Now,
		"username":   gofakeit.Username,
		"name":       gofakeit.Name,
		"email":      gofakeit.Email,
		"phone":      gofakeit.Phone,
		"int":        gofakeit.Int64,
		"digit":      gofakeit.Digit,
		"digitN":     gofakeit.DigitN,
		"letter":     gofakeit.Letter,
		"letterN":    gofakeit.LetterN,
		"word":       gofakeit.Word,
		"phrase":     gofakeit.Phrase,
		"loremIpsum": gofakeit.LoremIpsumSentence,
	}

	t, err := template.New("").Funcs(funcMap).Parse(text)
	if err != nil {
		return nil, fmt.Errorf("parse template: %w", err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, nil); err != nil {
		return nil, fmt.Errorf("execute template: %w", err)
	}

	return tpl.Bytes(), nil
}

func paramsValue(params map[string]string) func(string) string {
	return func(key string) string {
		return params[key]
	}
}
