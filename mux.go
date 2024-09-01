package gox

import (
	"net/http"

	m "github.com/narsus81/gox/internal/model"
)

func (g *Gox) HandleFunc(name string, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	// Add the route to the patterns map
	r := m.Route{Name: name, Handler: handler}
	g.patterns[pattern] = r
	// Wrap the handler with the default middleware
	g.mux.Handle(pattern, chainingMiddleware(http.HandlerFunc(handler), g.Chain...))
}

func (g *Gox) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.mux.ServeHTTP(w, r)
}
